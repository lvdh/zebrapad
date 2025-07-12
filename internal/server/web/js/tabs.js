// js/tabs.js
const tabCard = document.getElementById("tab-card");
const tabZpl = document.getElementById("tab-zpl");
const contentCard = document.getElementById("tab-content-card");
const contentZpl = document.getElementById("tab-content-zpl");
// default tab
tabCard.classList.add("is-active");
tabZpl.classList.remove("is-active");
contentCard.style.display = "";
contentZpl.style.display = "none";
// tab 'Card'
tabCard.addEventListener("click", function () {
  tabCard.classList.add("is-active");
  tabZpl.classList.remove("is-active");
  contentCard.style.display = "";
  contentZpl.style.display = "none";
});
// tab 'ZPL'
tabZpl.addEventListener("click", function () {
  tabZpl.classList.add("is-active");
  tabCard.classList.remove("is-active");
  contentZpl.style.display = "";
  contentCard.style.display = "none";
});
