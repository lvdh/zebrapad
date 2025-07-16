// js/forms.js

// functions
async function requestPreview(zpl) {
  const res = await fetch(
    "http://api.labelary.com/v1/printers/8dpmm/labels/2.25x0.75/0/",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: JSON.stringify(zpl),
    },
  );
  return res;
}

async function displayPreview(res, outputElId) {
  const resBlob = await res.blob();
  const img = document.createElement("img");
  img.src = URL.createObjectURL(resBlob);
  const outputDiv = document.getElementById(outputElId);
  while (outputDiv.firstChild) {
    // remove previous labels, if any
    outputDiv.removeChild(outputDiv.firstChild);
  }
  outputDiv.append(img);
  outputDiv.classList.remove("is-hidden");
}

async function displayPreviewNotification(res, resultElId) {
  const resultDiv = document.getElementById(resultElId);
  if (res.ok) {
    resultDiv.className = "notification is-info";
    resultDiv.textContent = "Label requested. Thank you, Labelary. 🤍";
  } else {
    const data = await res.json();
    resultDiv.className = "notification is-danger";
    resultDiv.textContent =
      data.error || "Error requesting label from Labelary. 🤕";
  }
  resultDiv.classList.remove("is-hidden");
}

async function printLabel(zpl) {
  const res = await fetch("/print", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ zpl }),
  });
  return res;
}

async function displayPrintNotification(res, resultElId) {
  const resultDiv = document.getElementById(resultElId);
  if (res.ok) {
    resultDiv.className = "notification is-success";
    resultDiv.textContent = "Label sent to printer. 🤞";
  } else {
    const data = await res.json();
    resultDiv.className = "notification is-danger";
    resultDiv.textContent = data.error || "Error printing label. 🖨️😱";
  }
  resultDiv.classList.remove("is-hidden");
}

// submit listeners for form 'Card'
document
  .getElementById("card-print")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const cardFullname = document.getElementById("card-fullname").value;
    const cardBarcode = document.getElementById("card-barcode").value;
    const zpl =
      "^XA^FO50,50^ADN,36,20^FD" + cardFullname + cardBarcode + "^FS^XZ";
    const res = await printLabel(zpl);
    await displayPrintNotification(res, "card-result");
  });
document
  .getElementById("card-preview")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const cardFullname = document.getElementById("card-fullname").value;
    const cardBarcode = document.getElementById("card-barcode").value;
    const zpl =
      "^XA^FO50,50^ADN,36,20^FD" + cardFullname + cardBarcode + "^FS^XZ";
    const res = await requestPreview(zpl);
    await displayPreviewNotification(res, "card-result");
    await displayPreview(res, "card-output");
  });

// submit listeners for form 'ZPL'
document
  .getElementById("zpl-print")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const zpl = document.getElementById("zpl-input").value;
    const res = await printLabel(zpl);
    displayPrintNotification(res, "zpl-result");
  });
document
  .getElementById("zpl-preview")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const zpl = document.getElementById("zpl-input").value;
    const res = await requestPreview(zpl);
    await displayPreviewNotification(res, "zpl-result");
    await displayPreview(res, "zpl-output");
  });
