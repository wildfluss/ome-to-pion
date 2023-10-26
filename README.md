# stream webrtc from OME to pion.ly 

How to repro (will try to receive from /app/fixme )

```bash
echo OME_HOST_IP=`ipconfig getifaddr en0` >> .env
```

```bash
docker compose up 
```

and start streaming to /app/fixme like this 

```bash
ffmpeg -re -i nice.mp3 -c:v libx264 -c:a aac -f flv rtmp://localhost:1935/app/fixme
```

and after a while it will signal , exchange offer answer and seemingly connect but after some few seconds pion drops connection it seems

```bash
ome-to-pion-webrtc-pion-1     | 2023/10/26 15:00:40 {"code":404,"error":"Cannot create offer"}
ome-to-pion-webrtc-pion-1     | 2023/10/26 15:00:40 Cannot create offer
ome-to-pion-ome-1             | [2023-10-26 15:00:41.057] I [SPRtcSig-t3333:10] Publisher | publisher.cpp:187  | Try to pull stream from local origin map: [#default#app/fixme]
ome-to-pion-ome-1             | [2023-10-26 15:00:41.057] I [SPRtcSig-t3333:10] Publisher | publisher.cpp:191  | Try to pull stream from origin map store: [#default#app/fixme]
ome-to-pion-ome-1             | [2023-10-26 15:00:41.057] E [SPRtcSig-t3333:10] Orchestrator | orchestrator.cpp:740  | Could not find Origin for the stream: [#default#app/fixme]
ome-to-pion-ome-1             | [2023-10-26 15:00:41.057] E [SPRtcSig-t3333:10] WebRTC Publisher | webrtc_publisher.cpp:333  | Cannot find stream (#default#app/fixme)
ome-to-pion-ome-1             | [2023-10-26 15:00:41.057] E [SPRtcSig-t3333:10] Signalling | rtc_signalling_server.cpp:407  | Cannot find stream [#default#app/fixme]
ome-to-pion-webrtc-pion-1     | 2023/10/26 15:00:41 {"code":404,"error":"Cannot create offer"}
ome-to-pion-webrtc-pion-1     | 2023/10/26 15:00:41 Cannot create offer
ome-to-pion-ome-1             | [2023-10-26 15:00:41.189] I [SPRTMP-t1935:26] MediaRouter | mediarouter_application.cpp:250  | [#default#app/fixme(100)] Trying to create a stream
ome-to-pion-ome-1             | [2023-10-26 15:00:41.189] I [SPRTMP-t1935:26] Monitor | application_metrics.cpp:57   | Create StreamMetrics(fixme/4314d152-cd58-467e-8c33-5acb5bbf9994/default/#default#app/fixme/i) for monitoring
ome-to-pion-ome-1             | [2023-10-26 15:00:41.189] I [SPRTMP-t1935:26] MediaRouter | mediarouter_application.cpp:348  | [#default#app/fixme(100)] Stream has been created
ome-to-pion-ome-1             | [2023-10-26 15:00:41.189] I [SPRTMP-t1935:26] Transcoder | transcoder_stream.cpp:46   | [#default#app/fixme(100)] Transcoder stream has been started
ome-to-pion-ome-1             | [2023-10-26 15:00:41.189] I [InboundWorker:27] MediaRouter | mediarouter_application.cpp:385  | [#default#app/fixme(100)] Stream has been prepared 
ome-to-pion-ome-1             | [Stream Info]
ome-to-pion-ome-1             | id(100), msid(0), output(fixme), SourceType(Rtmp), RepresentationType(Source), Created Time (Thu Oct 26 15:00:29 2023) UUID(4314d152-cd58-467e-8c33-5acb5bbf9994/default/#default#app/fixme/i)
ome-to-pion-ome-1             | 
ome-to-pion-ome-1             | 	Audio Track #1: Public Name(Audio_1) Variant Name(Audio) Bitrate(67.00Kb) Codec(6,AAC,Auto) BSF(AAC_RAW) Samplerate(22.1K) Format(s16, 16) Channel(mono, 1) timebase(1/1000)
ome-to-pion-ome-1             | 	Data  Track #2: Public Name(Data_2) Variant Name(Data) Codec(0,Unknown,Auto) BSF(ID3v2) timebase(1/1000)
ome-to-pion-ome-1             | [2023-10-26 15:00:41.190] I [InboundWorker:27] Transcoder | transcoder_stream.cpp:348  | [#default#app/fixme(100)] Output stream has been created. [#default#app/fixme(714050206)]
ome-to-pion-ome-1             | 
ome-to-pion-ome-1             | [2023-10-26 15:00:41.193] I [InboundWorker:27] MediaRouter | mediarouter_application.cpp:250  | [#default#app/fixme(714050206)] Trying to create a stream
ome-to-pion-ome-1             | [2023-10-26 15:00:41.193] I [InboundWorker:27] Monitor | application_metrics.cpp:57   | Create StreamMetrics(fixme/4314d152-cd58-467e-8c33-5acb5bbf9994/default/#default#app/fixme/o) for monitoring
ome-to-pion-ome-1             | [2023-10-26 15:00:41.193] I [InboundWorker:27] MediaRouter | mediarouter_application.cpp:348  | [#default#app/fixme(714050206)] Stream has been created
ome-to-pion-ome-1             | [2023-10-26 15:00:41.193] I [InboundWorker:27] Transcoder | transcoder_stream.cpp:94   | [#default#app/fixme(100)] Transcoder stream has been prepared
ome-to-pion-ome-1             | [2023-10-26 15:00:41.194] I [Decaac:36] Transcoder | decoder_aac.cpp:203  | [#default#app/fixme(100)] input track information: [audio] aac (LC), 22050 Hz, mono(1), fltp, 4 kbps, timebase: 1/1000, 
ome-to-pion-ome-1             | [2023-10-26 15:00:41.194] I [Decaac:36] Transcoder | transcoder_encoder.cpp:72   | [#1] hardware acceleration of the encoder is disabled. The library to be used will be Auto
ome-to-pion-ome-1             | [2023-10-26 15:00:41.195] I [Decaac:36] Transcoder | filter_resampler.cpp:147  | Resampler is enabled for track #1 using parameters. input: time_base=1/1000:sample_rate=22050:sample_fmt=fltp:channel_layout=mono / outputs: asettb=1/48000,aresample=async=1000,aresample=48000,aformat=sample_fmts=s16:channel_layouts=stereo,asetnsamples=n=960
ome-to-pion-ome-1             | [2023-10-26 15:00:42.234] I [OutboundWorker:28] MediaRouter | mediarouter_application.cpp:385  | [#default#app/fixme(714050206)] Stream has been prepared 
ome-to-pion-ome-1             | [Stream Info]
ome-to-pion-ome-1             | id(714050206), msid(0), output(fixme), SourceType(Transcoder), RepresentationType(Source), Created Time (Thu Oct 26 15:00:41 2023) UUID(4314d152-cd58-467e-8c33-5acb5bbf9994/default/#default#app/fixme/o)
ome-to-pion-ome-1             | 	>> Origin Stream Info
ome-to-pion-ome-1             | 	id(100), output(fixme), SourceType(Rtmp), Created Time (Thu Oct 26 15:00:29 2023)
ome-to-pion-ome-1             | 
ome-to-pion-ome-1             | 	Audio Track #0: Public Name(Audio_1) Variant Name(aac_audio) Bitrate(128.00Kb) Codec(6,AAC,Passthrough) BSF(AAC_RAW) Samplerate(22.1K) Format(fltp, 32) Channel(mono, 1) timebase(1/1000)
ome-to-pion-ome-1             | 	Audio Track #1: Public Name(Audio_1) Variant Name(opus_audio) Bitrate(128.00Kb) Codec(8,OPUS,libopus) BSF(OPUS) Samplerate(48.0K) Format(s16, 16) Channel(stereo, 2) timebase(1/48000)
ome-to-pion-ome-1             | 	Data  Track #2: Public Name(Data_2) Variant Name(Data) Codec(0,Unknown,Passthrough) BSF(ID3v2) timebase(1/1000)
ome-to-pion-ome-1             | 
ome-to-pion-ome-1             | [2023-10-26 15:00:42.236] I [OutboundWorker:28] WebRTC Publisher | rtc_stream.cpp:163  | RtcStream(#default#app/fixme) - Ignore unsupported codec(AAC)
ome-to-pion-ome-1             | [2023-10-26 15:00:42.236] I [OutboundWorker:28] WebRTC Publisher | rtc_stream.cpp:200  | WebRTC Stream has been created : fixme/714050206
ome-to-pion-ome-1             | Rtx(false) Ulpfec(false) JitterBuffer(false) PlayoutDelay(false min:0 max: 0)
ome-to-pion-ome-1             | [2023-10-26 15:00:42.237] I [OutboundWorker:28] Publisher | stream.cpp:212  | WebRTC Publisher Application application has started [fixme(714050206)] stream (MSID : 0)
ome-to-pion-ome-1             | [2023-10-26 15:00:42.237] I [OutboundWorker:28] Publisher | stream.cpp:212  | FilePublisher Application application has started [fixme(714050206)] stream (MSID : 0)
ome-to-pion-webrtc-pion-1     | 2023/10/26 15:00:42 {"candidates":[{"candidate":"candidate:0913826547 1 UDP 50 172.23.0.3 10004 typ host","sdpMLineIndex":0},{"candidate":"candidate:1349286507 1 UDP 50 185.233.95.134 10004 typ host","sdpMLineIndex":0}],"code":200,"command":"offer","iceServers":[{"credential":"airen","urls":["turn:185.233.95.134:3478?transport=tcp","turn:172.23.0.3:3478?transport=tcp"],"username":"ome"}],"ice_servers":[{"credential":"airen","urls":["turn:185.233.95.134:3478?transport=tcp","turn:172.23.0.3:3478?transport=tcp"],"user_name":"ome"}],"id":1708388794,"peer_id":0,"sdp":{"sdp":"v=0\r\no=OvenMediaEngine 100 2 IN IP4 127.0.0.1\r\ns=-\r\nt=0 0\r\na=group:BUNDLE Vq02eY\r\na=group:LS Vq02eY\r\na=msid-semantic:WMS IrBhuSF1ROULGKo8MqJ54i02e9XHVW3dZzNw\r\na=fingerprint:sha-256 AC:F8:3A:CA:79:1F:1C:94:60:60:6F:AF:4A:23:86:B7:45:FD:02:35:0F:31:07:B3:83:7D:CF:8D:AB:54:88:64\r\na=ice-options:trickle\r\na=ice-ufrag:qj2PGt\r\na=ice-pwd:bEfDepNi1BaK8odwTjI7kzP43mRLOUVM\r\nm=audio 9 UDP/TLS/RTP/SAVPF 110\r\nc=IN IP4 0.0.0.0\r\na=sendonly\r\na=mid:Vq02eY\r\na=setup:actpass\r\na=rtcp-mux\r\na=rtcp-rsize\r\na=msid:IrBhuSF1ROULGKo8MqJ54i02e9XHVW3dZzNw s9RlXN5tmfaEgjOdprAV4PkD7FKUGQcxeLCT\r\na=extmap:4 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time\r\na=rtpmap:110 OPUS/48000/2\r\na=fmtp:110 sprop-stereo=1;stereo=1;minptime=10;useinbandfec=1\r\na=ssrc:167944585 cname:oSpKZGs6wDiVrhXH\r\na=ssrc:167944585 msid:IrBhuSF1ROULGKo8MqJ54i02e9XHVW3dZzNw s9RlXN5tmfaEgjOdprAV4PkD7FKUGQcxeLCT\r\na=ssrc:167944585 mslabel:IrBhuSF1ROULGKo8MqJ54i02e9XHVW3dZzNw\r\na=ssrc:167944585 label:s9RlXN5tmfaEgjOdprAV4PkD7FKUGQcxeLCT\r\n","type":"offer"}}
ome-to-pion-webrtc-pion-1     | 2023/10/26 15:00:42 &{[{candidate:0913826547 1 UDP 50 172.23.0.3 10004 typ host 0} {candidate:1349286507 1 UDP 50 185.233.95.134 10004 typ host 0}] offer [{airen [turn:185.233.95.134:3478?transport=tcp turn:172.23.0.3:3478?transport=tcp] ome}] [{airen [turn:185.233.95.134:3478?transport=tcp turn:172.23.0.3:3478?transport=tcp] ome}] 1708388794 0 {offer v=0
ome-to-pion-webrtc-pion-1     | o=OvenMediaEngine 100 2 IN IP4 127.0.0.1
ome-to-pion-webrtc-pion-1     | s=-
ome-to-pion-webrtc-pion-1     | t=0 0
ome-to-pion-webrtc-pion-1     | a=group:BUNDLE Vq02eY
ome-to-pion-webrtc-pion-1     | a=group:LS Vq02eY
ome-to-pion-webrtc-pion-1     | a=msid-semantic:WMS IrBhuSF1ROULGKo8MqJ54i02e9XHVW3dZzNw
ome-to-pion-webrtc-pion-1     | a=fingerprint:sha-256 AC:F8:3A:CA:79:1F:1C:94:60:60:6F:AF:4A:23:86:B7:45:FD:02:35:0F:31:07:B3:83:7D:CF:8D:AB:54:88:64
ome-to-pion-webrtc-pion-1     | a=ice-options:trickle
ome-to-pion-webrtc-pion-1     | a=ice-ufrag:qj2PGt
ome-to-pion-webrtc-pion-1     | a=ice-pwd:bEfDepNi1BaK8odwTjI7kzP43mRLOUVM
ome-to-pion-webrtc-pion-1     | m=audio 9 UDP/TLS/RTP/SAVPF 110
ome-to-pion-webrtc-pion-1     | c=IN IP4 0.0.0.0
ome-to-pion-webrtc-pion-1     | a=sendonly
ome-to-pion-webrtc-pion-1     | a=mid:Vq02eY
ome-to-pion-webrtc-pion-1     | a=setup:actpass
ome-to-pion-webrtc-pion-1     | a=rtcp-mux
ome-to-pion-webrtc-pion-1     | a=rtcp-rsize
ome-to-pion-webrtc-pion-1     | a=msid:IrBhuSF1ROULGKo8MqJ54i02e9XHVW3dZzNw s9RlXN5tmfaEgjOdprAV4PkD7FKUGQcxeLCT
ome-to-pion-webrtc-pion-1     | a=extmap:4 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time
ome-to-pion-webrtc-pion-1     | a=rtpmap:110 OPUS/48000/2
ome-to-pion-webrtc-pion-1     | a=fmtp:110 sprop-stereo=1;stereo=1;minptime=10;useinbandfec=1
ome-to-pion-webrtc-pion-1     | a=ssrc:167944585 cname:oSpKZGs6wDiVrhXH
ome-to-pion-webrtc-pion-1     | a=ssrc:167944585 msid:IrBhuSF1ROULGKo8MqJ54i02e9XHVW3dZzNw s9RlXN5tmfaEgjOdprAV4PkD7FKUGQcxeLCT
ome-to-pion-webrtc-pion-1     | a=ssrc:167944585 mslabel:IrBhuSF1ROULGKo8MqJ54i02e9XHVW3dZzNw
ome-to-pion-webrtc-pion-1     | a=ssrc:167944585 label:s9RlXN5tmfaEgjOdprAV4PkD7FKUGQcxeLCT
ome-to-pion-webrtc-pion-1     |  <nil>}}
ome-to-pion-webrtc-pion-1     | 2023/10/26 15:00:42 ICEServer:{URLs:[turn:185.233.95.134:3478?transport=tcp turn:172.23.0.3:3478?transport=tcp] Username:ome Credential:airen CredentialType:password}
ome-to-pion-webrtc-pion-1     | 2023/10/26 15:00:42 Created a new RTCPeerConnection
ome-to-pion-webrtc-pion-1     | Peer Connection State has changed: connecting
ome-to-pion-webrtc-pion-1     | 2023/10/26 15:00:42 {answer v=0
ome-to-pion-webrtc-pion-1     | o=- 5322354730313924996 1698332442 IN IP4 0.0.0.0
ome-to-pion-webrtc-pion-1     | s=-
ome-to-pion-webrtc-pion-1     | t=0 0
ome-to-pion-webrtc-pion-1     | a=fingerprint:sha-256 E4:33:01:E3:B3:F2:C0:28:45:07:24:55:F1:EB:02:6F:C6:AC:40:90:AB:17:08:B3:91:D2:ED:9A:26:EE:EA:24
ome-to-pion-webrtc-pion-1     | a=group:BUNDLE Vq02eY
ome-to-pion-webrtc-pion-1     | m=audio 9 UDP/TLS/RTP/SAVPF 110
ome-to-pion-webrtc-pion-1     | c=IN IP4 0.0.0.0
ome-to-pion-webrtc-pion-1     | a=setup:active
ome-to-pion-webrtc-pion-1     | a=mid:Vq02eY
ome-to-pion-webrtc-pion-1     | a=ice-ufrag:ghrTrwNEJpQijbMf
ome-to-pion-webrtc-pion-1     | a=ice-pwd:ySByuDoyFYUpKeDNzMyJGJvDlZnJsEPu
ome-to-pion-webrtc-pion-1     | a=rtcp-mux
ome-to-pion-ome-1             | [2023-10-26 15:00:42.273] I [SPICE-t3478:19] ICE | ice_port.cpp:606  | Turn client has connected : <ClientSocket: 0xffff8f907010, #36, Connected, TCP, Nonblocking, 172.23.0.4:50068>
ome-to-pion-webrtc-pion-1     | a=rtcp-rsize
ome-to-pion-webrtc-pion-1     | a=rtpmap:110 OPUS/48000/2
ome-to-pion-webrtc-pion-1     | a=fmtp:110 sprop-stereo=1;stereo=1;minptime=10;useinbandfec=1
ome-to-pion-webrtc-pion-1     | a=ssrc:4214873356 cname:xKYwVRBWYZcamQGj
ome-to-pion-webrtc-pion-1     | a=ssrc:4214873356 msid:xKYwVRBWYZcamQGj xcScYEBnFocxxCZL
ome-to-pion-webrtc-pion-1     | a=ssrc:4214873356 mslabel:xKYwVRBWYZcamQGj
ome-to-pion-webrtc-pion-1     | a=ssrc:4214873356 label:xcScYEBnFocxxCZL
ome-to-pion-webrtc-pion-1     | a=msid:xKYwVRBWYZcamQGj xcScYEBnFocxxCZL
ome-to-pion-webrtc-pion-1     | a=sendrecv
ome-to-pion-webrtc-pion-1     |  0x40002586c0}
ome-to-pion-webrtc-pion-1     | 2023/10/26 15:00:42 oMEAnswer json: {"id":1708388794,"command":"answer","sdp":{"type":"answer","sdp":"v=0\r\no=- 5322354730313924996 1698332442 IN IP4 0.0.0.0\r\ns=-\r\nt=0 0\r\na=fingerprint:sha-256 E4:33:01:E3:B3:F2:C0:28:45:07:24:55:F1:EB:02:6F:C6:AC:40:90:AB:17:08:B3:91:D2:ED:9A:26:EE:EA:24\r\na=group:BUNDLE Vq02eY\r\nm=audio 9 UDP/TLS/RTP/SAVPF 110\r\nc=IN IP4 0.0.0.0\r\na=setup:active\r\na=mid:Vq02eY\r\na=ice-ufrag:ghrTrwNEJpQijbMf\r\na=ice-pwd:ySByuDoyFYUpKeDNzMyJGJvDlZnJsEPu\r\na=rtcp-mux\r\na=rtcp-rsize\r\na=rtpmap:110 OPUS/48000/2\r\na=fmtp:110 sprop-stereo=1;stereo=1;minptime=10;useinbandfec=1\r\na=ssrc:4214873356 cname:xKYwVRBWYZcamQGj\r\na=ssrc:4214873356 msid:xKYwVRBWYZcamQGj xcScYEBnFocxxCZL\r\na=ssrc:4214873356 mslabel:xKYwVRBWYZcamQGj\r\na=ssrc:4214873356 label:xcScYEBnFocxxCZL\r\na=msid:xKYwVRBWYZcamQGj xcScYEBnFocxxCZL\r\na=sendrecv\r\n"},"candidates":[]}
ome-to-pion-ome-1             | [2023-10-26 15:00:42.276] I [SPRtcSig-t3333:10] Monitor | stream_metrics.cpp:114  | A new session has started playing #default#app/fixme on the WebRTC publisher. WebRTC(1)/Stream total(1)/App total(1)
ome-to-pion-ome-1             | [2023-10-26 15:00:42.276] I [SPRtcSig-t3333:10] ICE | ice_port.cpp:368  | Added session: 0 (ufrag: qj2PGt:ghrTrwNEJpQijbMf)
ome-to-pion-webrtc-pion-1     | 2023/10/26 15:00:47 Blocking forever
ome-to-pion-ome-1             | [2023-10-26 15:01:12.408] I [SPICE-t3478:19] ICE | ice_port.cpp:620  | Turn client has disconnected : <ClientSocket: 0xffff8f907010, #36, Disconnected, TCP, Nonblocking, 172.23.0.4:50068>
ome-to-pion-webrtc-pion-1     | Peer Connection State has changed: failed
ome-to-pion-webrtc-pion-1     | Peer Connection has gone to failed exiting
ome-to-pion-ome-1             | [2023-10-26 15:01:12.420] I [DQICETmout:13] ICE | ice_port.cpp:436  | Removed session(0) from ICEPort | ice_seesions_with_id count(0) ice_sessions_with_ufrag(0) ice_sessions_with_address_pair(0) 
ome-to-pion-ome-1             | [2023-10-26 15:01:12.420] I [DQICETmout:13] WebRTC Publisher | webrtc_publisher.cpp:649  | IcePort is disconnected. : (#default#app/fixme/100) reason(6)
ome-to-pion-ome-1             | [2023-10-26 15:01:12.420] I [DQICETmout:13] WebRTC Publisher | webrtc_publisher.cpp:551  | Stop command received : #default#app/fixme/100
ome-to-pion-ome-1             | [2023-10-26 15:01:12.420] W [DQICETmout:13] ICE | ice_port.cpp:527  | Agent [Unknow, 0] has expired
ome-to-pion-ome-1             | [2023-10-26 15:01:12.421] I [DQICETmout:13] Monitor | stream_metrics.cpp:136  | A session has been stopped playing #default#app/fixme on the WebRTC publisher. Concurrent Viewers[WebRTC(0)/Stream total(0)/App total(0)]
ome-to-pion-ome-1             | [2023-10-26 15:01:12.421] I [DQICETmout:13] Signalling | rtc_signalling_server.cpp:450  | Client is disconnected: HttpConnection(0xffff8c801010) : WebSocket <ClientSocket: 0xffff8f840a10, #34, Connected, TCP, Nonblocking, 172.23.0.4:58704> TLS(Disabled) (#default#app / fixme, ufrag: local: qj2PGt, remote: ghrTrwNEJpQijbMf)
ome-to-pion-ome-1             | [2023-10-26 15:01:12.504] I [SPRtcSig-t3333:10] HTTP.Server | http_server.cpp:216  | Client(<ClientSocket: 0xffff8f840a10, #34, Closed, TCP, Nonblocking, 172.23.0.4:58704>) has been disconnected by *:3333

```

## Testing using ovenplayer

NB: nice.mp3 is a Q2 2023 Earnings Conference Call recording from here https://www.nice.com/company/investors/recorded-calls-archive 

Stream nice.mp3 , ~~FIXME? speed= 121x~~ use `-re `

```bash
$ ffmpeg -re -i nice.mp3 -c:v libx264 -c:a aac -f flv rtmp://localhost:1935/app/fixme
...
size=   35024kB time=01:06:28.95 bitrate=  71.9kbits/s speed=0.998x
...
```

Playback https://airensoft.gitbook.io/ovenmediaengine/streaming/webrtc-publishing#playback

and try ws://localhost:3333/app/fixme?transport=tcp in http://localhost:8090/ (use ../docker-compose.yaml whih has ovenplayer already ) maybe like this https://airensoft.gitbook.io/ovenmediaengine/streaming/webrtc-publishing#webrtc-over-tcp-with-ovenplayer

this also works ws://localhost:3333/app/fixme , but the sound is off as if its too fast 

## Output from OvenPlayer successfuly playing back ffmpeg stream in docker compose 

```bash
ome-to-pion-ome-1             | [2023-10-25 16:45:26.411] I [SPRtcSig-t3333:10] Publisher | publisher.cpp:191  | Try to pull stream from origin map store: [#default#app/fixme]
ome-to-pion-ome-1             | [2023-10-25 16:45:26.411] E [SPRtcSig-t3333:10] Orchestrator | orchestrator.cpp:740  | Could not find Origin for the stream: [#default#app/fixme]
ome-to-pion-ome-1             | [2023-10-25 16:45:26.411] E [SPRtcSig-t3333:10] WebRTC Publisher | webrtc_publisher.cpp:333  | Cannot find stream (#default#app/fixme)
ome-to-pion-ome-1             | [2023-10-25 16:45:26.411] E [SPRtcSig-t3333:10] Signalling | rtc_signalling_server.cpp:407  | Cannot find stream [#default#app/fixme]
ome-to-pion-ome-1             | [2023-10-25 16:45:26.413] I [SPRtcSig-t3333:10] Signalling | rtc_signalling_server.cpp:450  | Client is disconnected: HttpConnection(0xffffb9a00c10) : WebSocket <ClientSocket: 0xffffbd240a10, #36, Connected, TCP, Nonblocking, 172.22.0.1:51298> TLS(Disabled) (#default#app / fixme, ufrag: local: (N/A), remote: (N/A))
ome-to-pion-ome-1             | [2023-10-25 16:45:26.413] I [SPRtcSig-t3333:10] HTTP.Server | http_server.cpp:216  | Client(<ClientSocket: 0xffffbd240a10, #36, Closed, TCP, Nonblocking, 172.22.0.1:51298>) has been disconnected by *:3333
ome-to-pion-ome-1             | [2023-10-25 16:45:26.746] I [SPRTMP-t1935:26] MediaRouter | mediarouter_application.cpp:250  | [#default#app/fixme(100)] Trying to create a stream
ome-to-pion-ome-1             | [2023-10-25 16:45:26.746] I [SPRTMP-t1935:26] Monitor | application_metrics.cpp:57   | Create StreamMetrics(fixme/1bc8f346-194c-4a8b-95d3-6d13f9c4e03d/default/#default#app/fixme/i) for monitoring
ome-to-pion-ome-1             | [2023-10-25 16:45:26.746] I [SPRTMP-t1935:26] MediaRouter | mediarouter_application.cpp:348  | [#default#app/fixme(100)] Stream has been created
ome-to-pion-ome-1             | [2023-10-25 16:45:26.747] I [SPRTMP-t1935:26] Transcoder | transcoder_stream.cpp:46   | [#default#app/fixme(100)] Transcoder stream has been started
ome-to-pion-ome-1             | [2023-10-25 16:45:26.747] I [InboundWorker:27] MediaRouter | mediarouter_application.cpp:385  | [#default#app/fixme(100)] Stream has been prepared 
ome-to-pion-ome-1             | [Stream Info]
ome-to-pion-ome-1             | id(100), msid(0), output(fixme), SourceType(Rtmp), RepresentationType(Source), Created Time (Wed Oct 25 16:45:15 2023) UUID(1bc8f346-194c-4a8b-95d3-6d13f9c4e03d/default/#default#app/fixme/i)
ome-to-pion-ome-1             | 
ome-to-pion-ome-1             | 	Audio Track #1: Public Name(Audio_1) Variant Name(Audio) Bitrate(67.00Kb) Codec(6,AAC,Auto) BSF(AAC_RAW) Samplerate(22.1K) Format(s16, 16) Channel(mono, 1) timebase(1/1000)
ome-to-pion-ome-1             | 	Data  Track #2: Public Name(Data_2) Variant Name(Data) Codec(0,Unknown,Auto) BSF(ID3v2) timebase(1/1000)
ome-to-pion-ome-1             | 
ome-to-pion-ome-1             | [2023-10-25 16:45:26.748] I [InboundWorker:27] Transcoder | transcoder_stream.cpp:348  | [#default#app/fixme(100)] Output stream has been created. [#default#app/fixme(1534386032)]
ome-to-pion-ome-1             | [2023-10-25 16:45:26.753] I [InboundWorker:27] MediaRouter | mediarouter_application.cpp:250  | [#default#app/fixme(1534386032)] Trying to create a stream
ome-to-pion-ome-1             | [2023-10-25 16:45:26.753] I [InboundWorker:27] Monitor | application_metrics.cpp:57   | Create StreamMetrics(fixme/1bc8f346-194c-4a8b-95d3-6d13f9c4e03d/default/#default#app/fixme/o) for monitoring
ome-to-pion-ome-1             | [2023-10-25 16:45:26.753] I [InboundWorker:27] MediaRouter | mediarouter_application.cpp:348  | [#default#app/fixme(1534386032)] Stream has been created
ome-to-pion-ome-1             | [2023-10-25 16:45:26.753] I [InboundWorker:27] Transcoder | transcoder_stream.cpp:94   | [#default#app/fixme(100)] Transcoder stream has been prepared
ome-to-pion-ome-1             | [2023-10-25 16:45:26.754] I [Decaac:37] Transcoder | decoder_aac.cpp:203  | [#default#app/fixme(100)] input track information: [audio] aac (LC), 22050 Hz, mono(1), fltp, 4 kbps, timebase: 1/1000, 
ome-to-pion-ome-1             | [2023-10-25 16:45:26.754] I [Decaac:37] Transcoder | transcoder_encoder.cpp:72   | [#1] hardware acceleration of the encoder is disabled. The library to be used will be Auto
ome-to-pion-ome-1             | [2023-10-25 16:45:26.755] I [Decaac:37] Transcoder | filter_resampler.cpp:147  | Resampler is enabled for track #1 using parameters. input: time_base=1/1000:sample_rate=22050:sample_fmt=fltp:channel_layout=mono / outputs: asettb=1/48000,aresample=async=1000,aresample=48000,aformat=sample_fmts=s16:channel_layouts=stereo,asetnsamples=n=960
ome-to-pion-ome-1             | [2023-10-25 16:45:27.790] I [OutboundWorker:28] MediaRouter | mediarouter_application.cpp:385  | [#default#app/fixme(1534386032)] Stream has been prepared 
ome-to-pion-ome-1             | [Stream Info]
ome-to-pion-ome-1             | id(1534386032), msid(0), output(fixme), SourceType(Transcoder), RepresentationType(Source), Created Time (Wed Oct 25 16:45:26 2023) UUID(1bc8f346-194c-4a8b-95d3-6d13f9c4e03d/default/#default#app/fixme/o)
ome-to-pion-ome-1             | 	>> Origin Stream Info
ome-to-pion-ome-1             | 	id(100), output(fixme), SourceType(Rtmp), Created Time (Wed Oct 25 16:45:15 2023)
ome-to-pion-ome-1             | 
ome-to-pion-ome-1             | 	Audio Track #0: Public Name(Audio_1) Variant Name(aac_audio) Bitrate(128.00Kb) Codec(6,AAC,Passthrough) BSF(AAC_RAW) Samplerate(22.1K) Format(fltp, 32) Channel(mono, 1) timebase(1/1000)
ome-to-pion-ome-1             | 	Audio Track #1: Public Name(Audio_1) Variant Name(opus_audio) Bitrate(128.00Kb) Codec(8,OPUS,libopus) BSF(OPUS) Samplerate(48.0K) Format(s16, 16) Channel(stereo, 2) timebase(1/48000)
ome-to-pion-ome-1             | 	Data  Track #2: Public Name(Data_2) Variant Name(Data) Codec(0,Unknown,Passthrough) BSF(ID3v2) timebase(1/1000)
ome-to-pion-ome-1             | 
ome-to-pion-ome-1             | [2023-10-25 16:45:27.792] I [OutboundWorker:28] WebRTC Publisher | rtc_stream.cpp:163  | RtcStream(#default#app/fixme) - Ignore unsupported codec(AAC)
ome-to-pion-ome-1             | [2023-10-25 16:45:27.792] I [OutboundWorker:28] WebRTC Publisher | rtc_stream.cpp:200  | WebRTC Stream has been created : fixme/1534386032
ome-to-pion-ome-1             | Rtx(false) Ulpfec(false) JitterBuffer(false) PlayoutDelay(false min:0 max: 0)
ome-to-pion-ome-1             | [2023-10-25 16:45:27.792] I [OutboundWorker:28] Publisher | stream.cpp:212  | WebRTC Publisher Application application has started [fixme(1534386032)] stream (MSID : 0)
ome-to-pion-ome-1             | [2023-10-25 16:45:27.792] I [OutboundWorker:28] Publisher | stream.cpp:212  | FilePublisher Application application has started [fixme(1534386032)] stream (MSID : 0)
ome-to-pion-ome-1             | [2023-10-25 16:45:28.427] I [SPRtcSig-t3333:10] HTTP.Server | http_server.cpp:149  | Client(<ClientSocket: 0xffffbd240a10, #36, Connected, TCP, Nonblocking, 172.22.0.1:51310>) is connected on *:3333
ome-to-pion-ome-1             | [2023-10-25 16:45:28.428] I [SPRtcSig-t3333:10] Signalling | rtc_signalling_server.cpp:316  | New client is connected: <ClientSocket: 0xffffbd240a10, #36, Connected, TCP, Nonblocking, 172.22.0.1:51310>
ome-to-pion-ovenplayerdemo-1  | 172.22.0.1 - - [25/Oct/2023:16:45:28 +0000] "GET /null HTTP/1.1" 404 555 "http://localhost:8090/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36" "-"
ome-to-pion-admission-api-1   | ┌ Info: str
ome-to-pion-admission-api-1   | └   str = "{\"client\":{\"address\":\"172.22.0.1\",\"port\":51310,\"user_agent\":\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36\"},\"request\":{\"direction\":\"outgoing\",\"protocol\":\"WebRTC\",\"status\":\"opening\",\"time\":\"2023-10-25T16:45:28.438+00:00\",\"url\":\"ws://localhost:3333/app/fixme\"}}"
ome-to-pion-ome-1             | [2023-10-25 16:45:28.441] I [SPRtcSig-t3333:10] AccessController | access_controller.cpp:257  | AdmissionWebhooks queried http://admission-api:9595/admission whether client 172.22.0.1:51310 could access ws://localhost:3333/app/fixme. (Result : Allow Elapsed : 3 ms)
ome-to-pion-ome-1             | [2023-10-25 16:45:28.464] I [SPRtcSig-t3333:10] Monitor | stream_metrics.cpp:114  | A new session has started playing #default#app/fixme on the WebRTC publisher. WebRTC(1)/Stream total(1)/App total(1)
ome-to-pion-ome-1             | [2023-10-25 16:45:28.464] I [SPRtcSig-t3333:10] ICE | ice_port.cpp:368  | Added session: 0 (ufrag: GX4Irl:lQfy)
ome-to-pion-ome-1             | [2023-10-25 16:45:28.525] I [SPICE-t3478:19] ICE | ice_port.cpp:606  | Turn client has connected : <ClientSocket: 0xffffbd3f1010, #37, Connected, TCP, Nonblocking, 172.22.0.1:40076>
ome-to-pion-ome-1             | [2023-10-25 16:45:28.526] I [SPICE-t3478:19] ICE | ice_port.cpp:606  | Turn client has connected : <ClientSocket: 0xffffbd3f1210, #38, Connected, TCP, Nonblocking, 172.22.0.1:40078>
ome-to-pion-ome-1             | [2023-10-25 16:45:28.529] W [SPICE-t3478:19] ICE | ice_port.cpp:690  | Could not find agent(<SocketAddressPair: 0xffffb37fd7c8, local: 172.22.0.4:3478, remote: 172.22.0.1:40076>) information. Dropping... [Packet type : DTLS GateType : SEND_INDICATION]
ome-to-pion-ome-1             | [2023-10-25 16:45:28.529] I [SPICE-t3478:19] ICE | ice_port.cpp:843  | Session 0 uses candidate: <SocketAddressPair: 0xffffb37fd7c8, local: 172.22.0.4:3478, remote: 172.22.0.1:40076>
```

## OME signalling 

WebRTC requires signalling to exchange Offer SDP and Answer SDP, but this part is not standardized https://airensoft.gitbook.io/ovenmediaengine/v/0.10.0/webrtc-publishing#signalling

```bash
$ wscat --connect ws://localhost:3333/app/fixme
Connected (press CTRL+C to quit)
> { "command" :"request_offer" } 
< {"code":404,"error":"Cannot create offer"}
> 
```

in ome:

```bash
ome-to-pion-ome-1             | [2023-10-25 14:50:12.392] I [SPRtcSig-t3333:10] Publisher | publisher.cpp:187  | Try to pull stream from local origin map: [#default#app/fixme]
ome-to-pion-ome-1             | [2023-10-25 14:50:12.392] I [SPRtcSig-t3333:10] Publisher | publisher.cpp:191  | Try to pull stream from origin map store: [#default#app/fixme]
ome-to-pion-ome-1             | [2023-10-25 14:50:12.392] E [SPRtcSig-t3333:10] Orchestrator | orchestrator.cpp:740  | Could not find Origin for the stream: [#default#app/fixme]
ome-to-pion-ome-1             | [2023-10-25 14:50:12.392] E [SPRtcSig-t3333:10] WebRTC Publisher | webrtc_publisher.cpp:333  | Cannot find stream (#default#app/fixme)
ome-to-pion-ome-1             | [2023-10-25 14:50:12.392] E [SPRtcSig-t3333:10] Signalling | rtc_signalling_server.cpp:407  | Cannot find stream [#default#app/fixme]
```

then, start streaming using ffmpeg and try again and eventually after Admission allows the client 

```bash
> { "command" :"request_offer" } 
< {"candidates":[{"candidate":"candidate:0253986714 1 UDP 50 172.22.0.4 10004 typ host","sdpMLineIndex":0},{"candidate":"candidate:9621380754 1 UDP 50 185.233.95.134 10004 typ host","sdpMLineIndex":0}],"code":200,"command":"offer","iceServers":[{"credential":"airen","urls":["turn:185.233.95.134:3478?transport=tcp","turn:172.22.0.4:3478?transport=tcp"],"username":"ome"}],"ice_servers":[{"credential":"airen","urls":["turn:185.233.95.134:3478?transport=tcp","turn:172.22.0.4:3478?transport=tcp"],"user_name":"ome"}],"id":2135555650,"peer_id":0,"sdp":{"sdp":"v=0\r\no=OvenMediaEngine 100 2 IN IP4 127.0.0.1\r\ns=-\r\nt=0 0\r\na=group:BUNDLE ufdniv\r\na=group:LS ufdniv\r\na=msid-semantic:WMS 0eh1B7q5yODsdalHwLnjAIkTuXQvxK6VfS98\r\na=fingerprint:sha-256 6F:F3:D6:3C:57:41:C4:4D:68:72:1F:4E:81:B7:E2:80:1C:AF:AF:70:99:78:6E:0D:1A:F7:E9:A7:9F:3B:A3:74\r\na=ice-options:trickle\r\na=ice-ufrag:sd3CGJ\r\na=ice-pwd:rLTIKcAoOemSGh7z5X1iEfUgsZp3b8Yj\r\nm=audio 9 UDP/TLS/RTP/SAVPF 110\r\nc=IN IP4 0.0.0.0\r\na=sendonly\r\na=mid:ufdniv\r\na=setup:actpass\r\na=rtcp-mux\r\na=rtcp-rsize\r\na=msid:0eh1B7q5yODsdalHwLnjAIkTuXQvxK6VfS98 inpfL4J5dwmOPs3zYIrSguEbVtkh0jM8xvTR\r\na=extmap:4 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time\r\na=rtpmap:110 OPUS/48000/2\r\na=fmtp:110 sprop-stereo=1;stereo=1;minptime=10;useinbandfec=1\r\na=ssrc:2712961427 cname:71QdLBkAK2UYJ4vE\r\na=ssrc:2712961427 msid:0eh1B7q5yODsdalHwLnjAIkTuXQvxK6VfS98 inpfL4J5dwmOPs3zYIrSguEbVtkh0jM8xvTR\r\na=ssrc:2712961427 mslabel:0eh1B7q5yODsdalHwLnjAIkTuXQvxK6VfS98\r\na=ssrc:2712961427 label:inpfL4J5dwmOPs3zYIrSguEbVtkh0jM8xvTR\r\n","type":"offer"}}
```

