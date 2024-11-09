const watchDurationInput = document.getElementById("watchDuration");

let watchStartTime;
let accumulatedTime = 0;

// Start the timer when the video plays
video.addEventListener("play", () => {
  watchStartTime = Date.now();
  console.log("Video started playing at: ", watchStartTime);
});

// Stop the timer when the video pauses or ends, and update accumulated time
const stopTimer = () => {
  if (watchStartTime) {
    accumulatedTime += (Date.now() - watchStartTime) / 1000; // in seconds
    watchStartTime = null;
    watchDurationInput.value = accumulatedTime.toFixed(2);
    console.log("Watch duration: ", accumulatedTime.toFixed(2));

    return accumulatedTime.toFixed(2);
  }
};

video.addEventListener("pause", stopTimer);
video.addEventListener("ended", stopTimer);

window.addEventListener("beforeunload", stopTimer);
