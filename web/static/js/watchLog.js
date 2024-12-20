const watchDurationInput = document.getElementById("watchDuration");
const getwatchLog = document.getElementById("duration");

let watchStartTime;
let accumulatedTime = 0;

// Start the timer when the video plays
video.addEventListener("play", () => {
  watchStartTime = Date.now();
});

// Stop the timer when the video pauses or ends, and update accumulated time
const stopTimer = () => {
  if (watchStartTime) {
    accumulatedTime += (Date.now() - watchStartTime) / 1000;
    watchDurationInput.value = accumulatedTime.toFixed(2);
    watchStartTime = null;

    htmx.trigger(watchDurationInput, "submit");
  }
};

video.addEventListener("pause", stopTimer);
video.addEventListener("ended", stopTimer);
window.addEventListener("beforeunload", stopTimer);

getwatchLog.addEventListener("htmx:afterSwap", function (event) {
  console.log("The watchlog is updated to: ", getwatchLog.innerHTML);
  video.currentTime = parseFloat(getwatchLog.innerHTML);
  video.play();
}
);
console.log("the LOAD watchlog is....", getwatchLog.innerHTML);
