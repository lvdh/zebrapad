// js/forms.js
// form 'Card'
document
  .getElementById("card-print")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const cardFullname = document.getElementById("card-fullname").value;
    const cardBarcode = document.getElementById("card-barcode").value;
    const zpl =
      "^XA^FO50,50^ADN,36,20^FD" + cardFullname + cardBarcode + "^FS^XZ";
    const res = await fetch("/print", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ zpl }),
    });
    const resultDiv = document.getElementById("card-result");
    if (res.ok) {
      resultDiv.className = "notification is-success";
      resultDiv.textContent = "Label sent to printer!";
    } else {
      const data = await res.json();
      resultDiv.className = "notification is-danger";
      resultDiv.textContent = data.error || "Error sending label";
    }
    resultDiv.classList.remove("is-hidden");
  });
document
  .getElementById("card-render")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const cardFullname = document.getElementById("card-fullname").value;
    const cardBarcode = document.getElementById("card-barcode").value;
    const zpl =
      "^XA^FO50,50^ADN,36,20^FD" + cardFullname + cardBarcode + "^FS^XZ";
    const res = await fetch("/render", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ zpl }),
    });
    const resultDiv = document.getElementById("card-result");
    if (res.ok) {
      resultDiv.className = "notification is-info";
      resultDiv.textContent = "Label rendered successfully";
    } else {
      const data = await res.json();
      resultDiv.className = "notification is-danger";
      resultDiv.textContent = data.error || "Error rendering label";
    }
    const outputDiv = document.getElementById("card-output");
    outputDiv.classList.remove("is-hidden");
  });
// form 'ZPL'
document
  .getElementById("zpl-print")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const zpl = document.getElementById("zpl-input").value;
    const res = await fetch("/print", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ zpl }),
    });
    const resultDiv = document.getElementById("zpl-result");
    if (res.ok) {
      resultDiv.className = "notification is-success";
      resultDiv.textContent = "Label sent to printer!";
    } else {
      const data = await res.json();
      resultDiv.className = "notification is-danger";
      resultDiv.textContent = data.error || "Error sending label";
    }
    resultDiv.classList.remove("is-hidden");
  });
document
  .getElementById("zpl-render")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const zpl = document.getElementById("zpl-input").value;
    const res = await fetch("/render", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ zpl }),
    });
    const resultDiv = document.getElementById("zpl-result");
    if (res.ok) {
      resultDiv.className = "notification is-info";
      resultDiv.textContent = "Label rendered successfully";
    } else {
      const data = await res.json();
      resultDiv.className = "notification is-danger";
      resultDiv.textContent = data.error || "Error rendering label";
    }
    const outputDiv = document.getElementById("zpl-output");
    outputDiv.classList.remove("is-hidden");
  });
