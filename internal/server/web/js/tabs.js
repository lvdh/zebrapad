// static/js/tabs.js

const tabCard = document.getElementById("tab-card");
const tabZpl = document.getElementById("tab-zpl");
const refCard = document.getElementById("tab-card-ref");
const refZpl = document.getElementById("tab-zpl-ref");
const contentCard = document.getElementById("tab-content-card");
const contentZpl = document.getElementById("tab-content-zpl");

// default tab
tabCard.classList.add("is-active");
tabZpl.classList.remove("is-active");
refCard.classList.add("has-background-primary-light");
refZpl.classList.remove("has-background-primary-light");
contentCard.style.display = "";
contentZpl.style.display = "none";

// listener for tab 'Card'
tabCard.addEventListener("click", function () {
  tabCard.classList.add("is-active");
  tabZpl.classList.remove("is-active");
  refCard.classList.add("has-background-primary-light");
  refZpl.classList.remove("has-background-primary-light");
  contentCard.style.display = "";
  contentZpl.style.display = "none";
});

// listener for tab 'ZPL'
tabZpl.addEventListener("click", function () {
  tabZpl.classList.add("is-active");
  tabCard.classList.remove("is-active");
  refZpl.classList.add("has-background-primary-light");
  refCard.classList.remove("has-background-primary-light");
  contentZpl.style.display = "";
  contentCard.style.display = "none";
});
