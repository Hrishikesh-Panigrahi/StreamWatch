const video = document.getElementById('video');
const playBtn = document.getElementById('playBtn');
const fullscreenBtn = document.getElementById('fullscreenBtn');
const playerContainer = document.getElementById('player-container');
// Play or pause video
playBtn.addEventListener('click', () => {
    if (video.paused) {
        video.play();
        playBtn.textContent = 'Pause';
    } else {
        video.pause();
        playBtn.textContent = 'Play';
    }
});

// Toggle fullscreen mode
fullscreenBtn.addEventListener('click', () => {
    playerContainer.classList.toggle('fullscreen');
});

// Auto-hide controls when not interacting
let timeout;
playerContainer.addEventListener('mousemove', () => {
    clearTimeout(timeout);
    document.getElementById('controls').style.opacity = '1';
    timeout = setTimeout(() => {
        document.getElementById('controls').style.opacity = '0';
    }, 3000);
});