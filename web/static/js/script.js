const video = document.getElementById('video');
const playBtn = document.getElementById('playBtn');
const fullscreenBtn = document.getElementById('fullscreenBtn');
const resolutionDropdown = document.getElementById('resolutionDropdown');
const playerContainer = document.getElementById('player-container');

// Play or pause video
playBtn.addEventListener('click', () => {
    if (video.paused) {
        video.play();
        playBtn.textContent = '⏸';
    } else {
        video.pause();
        playBtn.textContent = '⏵';
    }
});

// Toggle fullscreen mode
fullscreenBtn.addEventListener('click', () => {
    playerContainer.classList.toggle('fullscreen');
});

// Handle resolution change
resolutionDropdown.addEventListener('change', () => {
    const selectedResolution = resolutionDropdown.value;

    // Update video source based on selected resolution
    const currentSource = video.querySelector('source');
    currentSource.src = `/stream/${selectedResolution}/${video.dataset.videoUuid}`;
    video.load();
    video.play();
});


if (Hls.isSupported()) {
    const hls = new Hls();
    hls.loadSource('/videos/{{ .data.Video.UUID }}/master.m3u8');
    hls.attachMedia(video);
} else if (video.canPlayType('application/vnd.apple.mpegurl')) {
    video.src = '/videos/{{ .data.Video.UUID }}/master.m3u8';
}