// static/js/tabs.js

const cardTab = document.getElementById("tab-card");
const cardRef = document.getElementById("tab-card-ref");
const cardContent = document.getElementById("tab-content-card");

const zplTab = document.getElementById("tab-zpl");
const zplRef = document.getElementById("tab-zpl-ref");
const zplContent = document.getElementById("tab-content-zpl");

const settingsTab = document.getElementById("tab-settings");
const settingsRef = document.getElementById("tab-settings-ref");
const settingsContent = document.getElementById("tab-content-settings");

// set default tab
cardTab.classList.add("is-active");
zplTab.classList.remove("is-active");
cardRef.classList.add("has-background-primary-light");
zplRef.classList.remove("has-background-primary-light");
cardContent.style.display = "";
zplContent.style.display = "none";

cardTab.addEventListener("click", function () {
  cardTab.classList.add("is-active");
  cardRef.classList.add("has-background-primary-light");
  cardContent.style.display = "";

  zplTab.classList.remove("is-active");
  zplRef.classList.remove("has-background-primary-light");
  zplContent.style.display = "none";

  settingsTab.classList.remove("is-active");
  settingsRef.classList.remove("has-background-primary-light");
  settingsContent.style.display = "none";
});

zplTab.addEventListener("click", function () {
  cardTab.classList.remove("is-active");
  cardRef.classList.remove("has-background-primary-light");
  cardContent.style.display = "none";

  zplTab.classList.add("is-active");
  zplRef.classList.add("has-background-primary-light");
  zplContent.style.display = "";

  settingsTab.classList.remove("is-active");
  settingsRef.classList.remove("has-background-primary-light");
  settingsContent.style.display = "none";
});

settingsTab.addEventListener("click", function () {
  cardTab.classList.remove("is-active");
  cardRef.classList.remove("has-background-primary-light");
  cardContent.style.display = "none";

  zplTab.classList.remove("is-active");
  zplRef.classList.remove("has-background-primary-light");
  zplContent.style.display = "none";

  settingsTab.classList.add("is-active");
  settingsRef.classList.add("has-background-primary-light");
  settingsContent.style.display = "";
});
