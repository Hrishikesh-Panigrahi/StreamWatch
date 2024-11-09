const watchDurationInput = document.getElementById("watchDuration");
const videoUUID = document.getElementById("UUID");

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
    accumulatedTime += (Date.now() - watchStartTime) / 1000; 
    watchDurationInput.value = accumulatedTime.toFixed(2);
    
    htmx.trigger(watchDurationInput, 'submit');
    watchStartTime = null;
  }
};

video.addEventListener("pause", stopTimer);
video.addEventListener("ended", stopTimer);

window.addEventListener("beforeunload", stopTimer);


// TODO: start the video where i left it off