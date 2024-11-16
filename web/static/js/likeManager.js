const like = document.querySelector("#likeCount");
const likeButton = document.querySelector("#send-like");

like.addEventListener("htmx:afterSwap", function (event) {
  console.log("The like count is updated to: ", like.innerHTML);
  //   check if the likbutton is already liked
  const isLiked =
    likeButton.classList.contains("text-green-600") &&
    likeButton.classList.contains("border-green-600");

  // Log if the button is liked or not
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
