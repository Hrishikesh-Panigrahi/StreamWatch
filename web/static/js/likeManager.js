const like = document.querySelector("#likeCount");
const likeButton = document.querySelector("#send-like");
var likeCount;
const VideoUUID = document.querySelector("#getID").value;

console.log(VideoUUID);

like.addEventListener("htmx:afterSwap", function (event) {
    console.log("The like count is updated to: ", like.innerHTML);
    likeCount = parseInt(like.innerHTML);
    console.log("Initial like count:", likeCount);
});

likeButton.addEventListener("click", () => {
    const isLiked = likeButton.classList.contains("text-green-600") &&
        likeButton.classList.contains("border-green-600");

    if (isLiked) {
        // Unlike action
        likeCount = likeCount - 1;
        likeButton.classList.remove("text-green-600", "border-green-600");
    } else {
        // Like action
        likeCount = likeCount + 1;
        likeButton.classList.add("text-green-600", "border-green-600");
    }

    like.innerHTML = likeCount;
    console.log("Updated like count:", likeCount);
});

window.addEventListener('beforeunload', () => {
    const data = {
        action: "leaving page",
        likeCount: likeCount,
    };
    const endpoint = `http://localhost:8080/video/${VideoUUID}/like`;

    navigator.sendBeacon(endpoint, JSON.stringify(data));
});


const dislike = document.querySelector("#dislikeCount");
const dislikeButton = document.querySelector("#send-dislike");
var dislikeCount;

dislike.addEventListener("htmx:afterSwap", function (event) {
    console.log("The dislike count is updated to: ", dislike.innerHTML);
    dislikeCount = parseInt(dislike.innerHTML);
    console.log("Initial dislike count:", dislikeCount);
});

dislikeButton.addEventListener("click", () => {
    const isLiked = dislikeButton.classList.contains("text-red-600") &&
        dislikeButton.classList.contains("border-red-600");

    if (isLiked) {
        // Unlike action
        dislikeCount = dislikeCount - 1;
        dislikeButton.classList.remove("text-red-600", "border-red-600");
    } else {
        // Like action
        dislikeCount = dislikeCount + 1;
        dislikeButton.classList.add("text-red-600", "border-red-600");
    }

    dislike.innerHTML = dislikeCount;
    console.log("Updated dislike count:", dislikeCount);
});

window.addEventListener('beforeunload', () => {
    const data = {
        action: "leaving page",
        dislikeCount: dislikeCount,
    };
    const endpoint = `http://localhost:8080/video/${VideoUUID}/dislike`;

    navigator.sendBeacon(endpoint, JSON.stringify(data));
});