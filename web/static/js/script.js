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
