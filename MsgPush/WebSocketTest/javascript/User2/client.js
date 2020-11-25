(function(win) {
    const rawHeaderLen = 16;
    const packetOffset = 0;
    const headerOffset = 4;
    const opOffset = 6;
    const seqOffset = 8;
    //const seqOffset = 12;

    var Client = function(options) {
        var MAX_CONNECT_TIMES = 10;
        var DELAY = 15000;
        this.options = options || {};
        this.createConnect(MAX_CONNECT_TIMES, DELAY);
    }

    var appendMsg = function(text) {
        var span = document.createElement("SPAN");
        var text = document.createTextNode(text);
        span.appendChild(text);
        document.getElementById("box").appendChild(span);
    }

    Client.prototype.createConnect = function(max, delay) {
        var self = this;
        if (max === 0) {
            return;
        }
        connect();

        var textDecoder = new TextDecoder();
        var textEncoder = new TextEncoder();
        var heartbeatInterval;
        var messageInterval;
        function connect() {
            //var ws = new WebSocket('ws://sh.tony.wiki:3102/sub');
            var ws = new WebSocket('ws://127.0.0.1:8081/ws/sport?userId=25ed060d-2ba7-11eb-b1a3-00163e2ed191&continue=1&sportId=b57b8bf5-2e60-11eb-a64c-00163e2ed191&sportType=1');
            ws.binaryType = 'arraybuffer';
            ws.onopen = function(evt) {
                appendMsg(evt)
                heartbeat();
                heartbeatInterval = setInterval(heartbeat, 30 * 1000);
                messageInterval = setInterval(sendMessage, 1 * 1000);
            }

            ws.onerror = function(evt){
                appendMsg(evt)
                appendMsg("连接错误")
            }
            ws.onmessage = function(evt) {
                var data = evt.data;
                var dataView = new DataView(data, 0);
                var op = dataView.getInt32(0);
                var seq = dataView.getInt32(4);
                var headLen = 8
                switch(op) {
                    case 9:
                        // auth reply ok
                        document.getElementById("status").innerHTML = "<color style='color:green'>ok<color>";
                        appendMsg("receive: auth reply");
                        // send a heartbeat to server
                       // heartbeat();
                        //heartbeatInterval = setInterval(heartbeat, 30 * 1000);
                        //messageInterval = setInterval(sendMessage, 1 * 1000);
                        break;
                    case 1:
                        // receive a heartbeat from server
                        console.log("receive: heartbeat");
                        appendMsg("receive: heartbeat reply");
                        break;
                    case 2:
                        // batch message

                        var type = dataView.getInt32(headLen)
                        var status = dataView.getInt32(headLen + 4)
                        var sportHeadLen = headLen + 8

                        var curDis = dataView.getInt32(sportHeadLen)
                        var curDisOffset = sportHeadLen + 4
                        var curPace = dataView.getInt32(curDisOffset)
                        var curPaceOffset = curDisOffset + 4
                        var curTime = dataView.getInt32(curPaceOffset)
                        var curTimeOffset = curPaceOffset + 4
                        var curCal = dataView.getInt32(curTimeOffset)
                        appendMsg("receive curDis = "+curDis + ",curPace:"+curPace + ",curTime:" + curTime + ",curCal:"+ curCal)
                        break;
                    case 3:
                        appendMsg("token验证失败")
                        if (heartbeatInterval) clearInterval(heartbeatInterval);
                        if (messageInterval) clearInterval(messageInterval);
                        break;

                    case 4:
                        var buflen = dataView.byteLength - headLen
                        var buf = new ArrayBuffer(buflen);
                        var bufView = new Uint8Array(buf);
                        for(var i = 0; i < buflen; i ++){
                            bufView[i] = dataView.getUint8(headLen + i)
                    }
                        appendMsg("验证成功，token:" + bufView.toString(16))
                    default:
                        var msgBody = textDecoder.decode(data.slice(headerLen, packetLen));
                        messageReceived(ver, msgBody);
                        appendMsg("receive: ver=" + ver + " op=" + op + " seq=" + seq + " message=" + msgBody);
                        break
                }
            }

            ws.onclose = function() {
                if (heartbeatInterval) clearInterval(heartbeatInterval);
                setTimeout(reConnect, delay);
                appendMsg("关闭连接")
                document.getElementById("status").innerHTML =  "<color style='color:red'>failed<color>";
            }

            var heartbeatSeq = 0;
            function heartbeat() {
                heartbeatSeq++;
                var headerBuf = new ArrayBuffer(8);
                var headerView = new DataView(headerBuf, 0);
                headerView.setInt32(0, 1);
                headerView.setInt32(4, heartbeatSeq);
                ws.send(headerBuf);
                console.log("send: heartbeat");
                appendMsg("send: heartbeat");
            }
            var seq = 1;
            function sendMessage(){
                var msg = "你好:"+seq;
                //var packLen = rawHeaderLen + msg.toString().length;

                var packLen = 32

                var packBuf = new ArrayBuffer(packLen);
                var packView = new DataView(packBuf, 0);

                packView.setInt32(0, 2); //op
                packView.setInt32(4, seq);      //seq
                packView.setInt32(8, 1); //type
                packView.setInt32(12, 1);//status
                packView.setInt32(16, seq + 1);     //curDis
                packView.setInt32(20, seq + 2);     //curPace
                packView.setInt32(24, seq + 3);     //curTime
                packView.setInt32(28, seq + 4);     //curCal
                ws.send(packBuf);
                seq += 1;
                appendMsg("send: Msg:" + seq);
            }

            function auth() {
                var body = "visus"
                var headerBuf = new ArrayBuffer(8);
                var dataView = new DataView(headerBuf, 0);
                var bodyBuf = textEncoder.encode(body);
                dataView.setInt32(packetOffset, 1); // op
                dataView.setInt32(4, 1); //seq

                ws.send(mergeArrayBuffer(headerBuf, bodyBuf));

                appendMsg("send: auth token: ");
            }

            function messageReceived(ver, body) {
                var notify = self.options.notify;
                if(notify) notify(body);
                console.log("messageReceived:", "ver=" + ver, "body=" + body);
            }

            function mergeArrayBuffer(ab1, ab2) {
                var u81 = new Uint8Array(ab1),
                    u82 = new Uint8Array(ab2),
                    res = new Uint8Array(ab1.byteLength + ab2.byteLength);
                appendMsg("ab1.toString:"+ ab1.toString() + " ab2.toString:" + ab2.toString())
                res.set(u81, 0);
                res.set(u82, ab1.byteLength);
                return res.buffer;
            }

            function char2ab(str) {
                var buf = new ArrayBuffer(str.length);
                var bufView = new Uint8Array(buf);
                for (var i=0; i<str.length; i++) {
                    bufView[i] = str[i];
                }
                return buf;
            }

        }

        function reConnect() {
            self.createConnect(--max, delay * 2);
        }
    }

    win['MyClient'] = Client;
})(window);
