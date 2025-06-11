import {useEffect, useRef, useState} from 'react';

function App() {
  const [socket, setSocket] = useState(null);
  const [roomId, setRoomId] = useState("");

  const localVideoRef = useRef(null);
  const remoteVideoRef = useRef(null);
  const peerConnection = useRef(null);

  const servers = {
    iceServers: [{
      urls: 'stun:stun.l.google.com:19302'  // Google's public STUN server
    }]
  }

  useEffect(() => {
    const ws = new WebSocket('ws://localhost:8080');
    setSocket(ws);

    ws.onmessage = async (event) => {
      const data = JSON.parse(event.data);
      if (data.type === 'offer') {
        await peerConnection.current.setRemoteDescription(data);
        const answer = await peerConnection.current.createAnswer();
        await peerConnection.current.setLocalDescription(answer);
        ws.send(JSON.stringify({
          type: "signal", roomId, signalData: peerConnection.current.localDescription,
        }));
      } else if (data.type === 'answer') {
        await peerConnection.current.setRemoteDescription(data);
      } else if (data.type === 'candidate') {
        await peerConnection.current.addIceCandidate(data.candidate);
      }
    };
    return () => {
      ws.close();
    }
  }, [roomId]);

  const joinRoom = async () => {
    if (!roomId.trim()) {
      alert('Please enter a room ID');
      return;
    }
    socket.send(JSON.stringify({type: 'join', roomId}));
    // Initialize local video stream
    const stream = await navigator.mediaDevices.getUserMedia({
      video: {frameRate: {ideal: 30, max: 60}},
    });

    localVideoRef.current.srcObject = stream;
    peerConnection.current = new RTCPeerConnection(servers);

    stream.getTracks().forEach(track => {
      peerConnection.current.addTrack(track, stream);
    });

    // ICE Candidate handling
    peerConnection.current.onicecandidate = (event) => {
      if (event.candidate) {
        socket.send(JSON.stringify({
          type: "signal", roomId: roomId, signalData: {type: "candidate", candidate: event.candidate}
        }));
      }
    };
    peerConnection.current.ontrack = (event) => {
      remoteVideoRef.current.srcObject = event.streams[0];
    };

    const offer = await peerConnection.current.createOffer();
    await peerConnection.current.setLocalDescription(offer);

    // SDP(Session Description Protocol) offer 생성
    socket.send(JSON.stringify({
      type: "signal", roomId, signalData: peerConnection.current.localDescription,
    }));
  }

  return (<div className="App">
    <h1>React WebRTC Video Chat</h1>
    <input
      type="text"
      placeholder="Enter Room ID"
      value={roomId}
      onChange={(e) => setRoomId(e.target.value)}
    />
    <button onClick={joinRoom}> Join Room</button>

    <div>
      <div>
        Local Video
        <video ref={localVideoRef} autoPlay></video>
      </div>

      <div>
        Remote Video
        <video ref={remoteVideoRef} autoPlay></video>
      </div>
    </div>
  </div>);
};

export default App;
