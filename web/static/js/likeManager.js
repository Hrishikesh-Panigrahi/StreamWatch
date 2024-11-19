const like = document.querySelector("#likeCount");
const likeButton = document.querySelector("#send-like");

like.addEventListener("htmx:afterSwap", function (event) {
  console.log("The like count is updated to: ", like.innerHTML);
  //   check if the likbutton is already liked
  const isLiked =
    likeButton.classList.contains("text-green-600") &&
    likeButton.classList.contains("border-green-600");
  
  if (isLiked) {
    likeButton.addEventListener("click", function () {
      console.log("The video is unliked.");
      like.innerHTML = parseInt(like.innerHTML) - 1;
      likeButton.classList.remove("text-green-600");
      likeButton.classList.remove("border-green-600");
    });
  } else {
    likeButton.addEventListener("click", function () {
      console.log("The video is liked.");
      like.innerHTML = parseInt(like.innerHTML) + 1;
      likeButton.classList.add("text-green-600");
      likeButton.classList.add("border-green-600");
    });
  }
});

const dislike = document.querySelector("#dislikeCount");
const dislikeButton = document.querySelector("#send-dislike");

dislike.addEventListener("htmx:afterSwap", function (event) {
  console.log("The dislike count is updated to: ", dislike.innerHTML);
  //   check if the dislike button is already disliked
  const isDisliked =
    dislikeButton.classList.contains("text-red-600") &&
    dislikeButton.classList.contains("border-red-600");

  if (isDisliked) {
    dislikeButton.addEventListener("click", function () {
      console.log("The video is undisliked.");
      dislike.innerHTML = parseInt(dislike.innerHTML) - 1;
      dislikeButton.classList.remove("text-red-600");
      dislikeButton.classList.remove("border-red-600");
    });
  } else {
    dislikeButton.addEventListener("click", function () {
      console.log("The video is disliked.");
      dislike.innerHTML = parseInt(dislike.innerHTML) + 1;
      dislikeButton.classList.add("text-red-600");
      dislikeButton.classList.add("border-red-600");
    });
  }
});

