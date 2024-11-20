const toggleSidebar = document.getElementById("toggleSidebar");
const closeSidebar = document.getElementById("closeSidebar");
const sidebar = document.getElementById("sidebar");

toggleSidebar.addEventListener("click", function () {
    if (sidebar.classList.contains("hidden")) {
        sidebar.classList.remove("hidden");
        sidebar.classList.add("translate-x-0");
    } else {
        sidebar.classList.add("hidden");
        sidebar.classList.remove("translate-x-0");
    }
});

closeSidebar.addEventListener("click", () => {
    sidebar.classList.add("hidden");
});