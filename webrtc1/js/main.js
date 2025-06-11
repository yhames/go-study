"use strict";

const localVideo = document.getElementById("localVideo");
const startButton = document.getElementById("startButton");
let localStream;
let localPeerConnection;

const remoteVideo = document.getElementById("remoteVideo");
const remoteButton = document.getElementById("remoteButton");
remoteButton.disabled = true;
let remoteStream;
let remotePeerConnection;

const cancelButton = document.getElementById("cancelButton");
cancelButton.disabled = true;

async function startAction() {
    startButton.disabled = true;
    const mediaStream = {
        video: {frameRate: {ideal: 30, max: 30}},
        audio: true,
    }
    localStream = await navigator.mediaDevices.getUserMedia(mediaStream);
    localVideo.srcObject = localStream;
    remoteButton.disabled = false;
}

async function remoteAction() {
    remoteButton.disabled = true;
    cancelButton.disabled = false;
    const videoTracks = localStream.getVideoTracks();
    if (videoTracks.length > 0) {
        console.log("Using video device: " + videoTracks[0].label);
    }
    const audioTracks = localStream.getAudioTracks();
    if (audioTracks.length > 0) {
        console.log("Using audio device: " + audioTracks[0].label);
    }

    const server = null;
    localPeerConnection = new RTCPeerConnection(server);    // default localhost
    localPeerConnection.addEventListener("icecandidate", handlerConnection);   // ICE 등록
    localPeerConnection.addEventListener("iceconnectionstatechange", handleConnectionChanged);   // ICE 상태 변경

    remotePeerConnection = new RTCPeerConnection(server);    // default localhost
    remotePeerConnection.addEventListener("icecandidate", handlerConnection);   // ICE 등록
    remotePeerConnection.addEventListener("iceconnectionstatechange", handleConnectionChanged);   // ICE 상태 변경

    remotePeerConnection.addEventListener("addstream", getRemoteMedia);
    localPeerConnection.addStream(localStream);

    // SDP(Session Description Protocol) 설정
    const offerOptions = {
        offerToReceiveVideo: true,
    }
    const sessionDescriptionInit = await localPeerConnection.createOffer(offerOptions);
    await afterOfferCreated(sessionDescriptionInit);
}

function cancelAction() {
    if (localPeerConnection) {
        localPeerConnection.close();
        localPeerConnection = null;
    }
    if (remotePeerConnection) {
        remotePeerConnection.close();
        remotePeerConnection = null;
    }
    cancelButton.disabled = true;
    remoteButton.disabled = false;
}

async function afterOfferCreated(sessionDescriptionInit) {
    // localPeerConnection에 SDP 설정
    // localPeer는 내가 전송하고자 하는 정보를 offer로 설정
    await localPeerConnection.setLocalDescription(sessionDescriptionInit);
    console.log("localPeerConnection established with SDP: " + localPeerConnection.localDescription.sdp);
    // remotePeerConnection에 localPeerConnection의 offer 설정
    await remotePeerConnection.setRemoteDescription(sessionDescriptionInit)
    console.log("remotePeerConnection established with SDP: " + remotePeerConnection.remoteDescription.sdp);

    // remotePeerConnection에 SDP 설정
    const answer = await remotePeerConnection.createAnswer();
    await remotePeerConnection.setLocalDescription(answer);
    console.log("remotePeerConnection answer created: " + remotePeerConnection.localDescription.sdp);
    // localPeerConnection에 remotePeerConnection의 answer 설정
    await localPeerConnection.setRemoteDescription(answer);
    console.log("localPeerConnection answer set: " + localPeerConnection.remoteDescription.sdp);
}

async function handlerConnection(event) {
    const target = event.target;
    const ice = event.candidate;
    if (ice) {
        const newIce = new RTCIceCandidate(ice);
        const otherPeer = getOtherPeer(target);
        await otherPeer.addIceCandidate(newIce);
        console.log("ICE candidate added: " + ice.candidate);
    }
}

function handleConnectionChanged(event) {
    console.log("ICE connection state changed: " + event.target);
}

function getOtherPeer(peer) {
    return (peer === localPeerConnection) ? remotePeerConnection : localPeerConnection;
}

function getRemoteMedia(event) {
    const stream = event.stream;
    remoteVideo.srcObject = stream;
    remoteStream = stream;
    console.log("Remote stream added: " + stream.id);
}

startButton.addEventListener("click", startAction);
remoteButton.addEventListener("click", remoteAction);
cancelButton.addEventListener("click", cancelAction);