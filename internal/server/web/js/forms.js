// js/forms.js
async function requestLabel(zpl) {
  const response = await fetch(
    `http://api.labelary.com/v1/printers/8dpmm/labels/2.25x0.75/0/`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: JSON.stringify(zpl),
    },
  );
  console.log(response);
  return response;
}

// form 'Card'
document
  .getElementById("card-print")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    // parse
    const cardFullname = document.getElementById("card-fullname").value;
    const cardBarcode = document.getElementById("card-barcode").value;
    const zpl =
      "^XA^FO50,50^ADN,36,20^FD" + cardFullname + cardBarcode + "^FS^XZ";
    // print
    const response = await fetch("/print", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ zpl }),
    });
    // notif
    const resultDiv = document.getElementById("card-result");
    if (response.ok) {
      resultDiv.className = "notification is-success";
      resultDiv.textContent = "Label sent to printer!";
    } else {
      const data = await response.json();
      resultDiv.className = "notification is-danger";
      resultDiv.textContent = data.error || "Error sending label";
    }
    resultDiv.classList.remove("is-hidden");
  });
document
  .getElementById("card-preview")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    // parse
    const cardFullname = document.getElementById("card-fullname").value;
    const cardBarcode = document.getElementById("card-barcode").value;
    const zpl =
      "^XA^FO50,50^ADN,36,20^FD" + cardFullname + cardBarcode + "^FS^XZ";
    // request label
    const response = await requestLabel(zpl);
    // display notif
    const resultDiv = document.getElementById("card-result");
    if (response.ok) {
      resultDiv.className = "notification is-success";
      resultDiv.textContent = "Label requested. Thank you, Labelary. 🤍";
    } else {
      const data = await response.json();
      resultDiv.className = "notification is-danger";
      resultDiv.textContent = data.error || "Error sending label";
    }
    resultDiv.classList.remove("is-hidden");
    // display label
    const responseBlob = await response.blob();
    const img = document.createElement("img");
    img.src = URL.createObjectURL(responseBlob);
    img.id = "card-label";
    const outputDiv = document.getElementById("card-output");
    while(outputDiv.firstChild) {
        // remove previous labels, if any
        outputDiv.removeChild(outputDiv.firstChild);
    }
    outputDiv.append(img);
    outputDiv.classList.remove("is-hidden");
  });
// form 'ZPL'
document
  .getElementById("zpl-print")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const zpl = document.getElementById("zpl-input").value;
    const response = await fetch("/print", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ zpl }),
    });
    const resultDiv = document.getElementById("zpl-result");
    if (response.ok) {
      resultDiv.className = "notification is-success";
      resultDiv.textContent = "Label sent to printer!";
    } else {
      const data = await response.json();
      resultDiv.className = "notification is-danger";
      resultDiv.textContent = data.error || "Error sending label";
    }
    resultDiv.classList.remove("is-hidden");
  });
document
  .getElementById("zpl-preview")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const zpl = document.getElementById("zpl-input").value;
    const response = await fetch("/preview", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ zpl }),
    });
    const resultDiv = document.getElementById("zpl-result");
    if (response.ok) {
      resultDiv.className = "notification is-info";
      resultDiv.textContent = "Label rendered successfully";
    } else {
      const data = await response.json();
      resultDiv.className = "notification is-danger";
      resultDiv.textContent = data.error || "Error rendering label";
    }
    const outputDiv = document.getElementById("zpl-output");
    outputDiv.classList.remove("is-hidden");
  });
