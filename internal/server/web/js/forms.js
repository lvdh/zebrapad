// static/js/forms.js

// functions
async function requestPreview(zpl) {
  // NOTE: .replace() works around escape shenanigans coming from textarea
  zplNormalized = JSON.stringify(zpl).replace("\\&", "\&");
  console.log(zplNormalized);
  const res = await fetch(
    // NOTE: label type is fixed, Zebra 800272-075:
    //   size in metric: 57.15mm x 19.05mm
    //   size in freedom units: 2.25" x 0.75"
    "http://api.labelary.com/v1/printers/8dpmm/labels/2.25x0.75/0/",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: zplNormalized,
    },
  );
  return res;
}

async function displayPreview(res, outputElId) {
  const resBlob = await res.blob();
  const img = document.createElement("img");
  img.src = URL.createObjectURL(resBlob);
  img.style.backgroundColor = "#000";
  img.style.padding = "1px";
  img.style.borderRadius = "20px";
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
    resultDiv.textContent = "Label sent to printer. 🖨️";
  } else {
    const data = await res.json();
    resultDiv.className = "notification is-danger";
    resultDiv.textContent = data.error || "Error printing label. 🖨️😱";
  }
  resultDiv.classList.remove("is-hidden");
}

function validateCardForm() {
  const cardBarcode = document.getElementById("card-barcode").value;
  const cardFullName = document.getElementById("card-fullname").value;
  const cardFullNameMsg = document.getElementById("card-fullname-msg");
  const cardBarcodeMsg = document.getElementById("card-barcode-msg");
  // clear previous error messages
  cardFullNameMsg.textContent = "";
  cardBarcodeMsg.textContent = "";
  // full name must be filled out and <=30 chars
  if (!cardFullName || cardFullName.length > 30) {
    let msg = "ℹ️ Please provide a name, 30 characters or less.";
    cardFullNameMsg.textContent = msg;
    throw new Error(msg);
  }
  // barcode must be filled out and exactly 9 digits
  if (!cardBarcode || cardBarcode.length != 9 || !parseInt(cardBarcode)) {
    let msg = "ℹ️ Please provide a barcode value of exactly 9 digits.";
    cardBarcodeMsg.textContent = msg;
    throw new Error(msg);
  }
  return {
    cardFullName: cardFullName,
    cardBarcode: cardBarcode,
  };
}

function validateZplForm() {
  const zpl = document.getElementById("zpl-input").value;
  zplInputMsg = document.getElementById("zpl-input-msg");
  zplInputMsg.textContent = "";
  if (!zpl.trim().startsWith("^XA") || !zpl.trim().endsWith("^XZ")) {
    let msg = "ℹ️ Please provide valid ZPL. See documentation to the right.";
    zplInputMsg.textContent = msg;
    throw new Error(msg);
  }
  return zpl;
}

function renderZplTemplate(formData) {
  const zplTemplate = `^XA
  
  ^FX-- BARCODE:^FS
  ^FX--   NOTE: MUST BE 9 DIGITS (FIXED POSITION)^FS
  ^FO78,10
  ^BY3
  ^BCN,70,,,,A
  ^FD999999999^FS
  
  ^FX-- FULL NAME:^FS
  ^FX--   NOTE: MAX. 30 CHARS^FS
  ^FO0,120
  ^FB457,,,C,
  ^A0,,30,
  ^FDBill the Lizard\\&^FS
  
  ^FX-- ASSUMED CONFIG:^FS
  ^FX--   LABEL:       Zebra 800272-075^FS
  ^FX--   PRINTER:     Zebra GK420t^FS
  ^FX--   DENSITY:     8 DPMM^FS
  ^FX-- LABEL:^FS
  ^FX--   RESOLUTION:  457.2 DOTS x 152.4 DOTS^FS
  ^FX-- BARCODE:^FS
  ^FX--   WIDTH:       297 DOTS^FS
  ^FX--     CALC:      (9x digit)*(11 dots/digit)*(BY=3)^FS
  ^FX--   SIDES:       80 DOTS^FS
  ^FX--     CALC:      ROUND((457.2-297)/2)^FS
  
  ^XZ`;
  return zplTemplate
    .replace("999999999", formData.cardBarcode)
    .replace("Bill the Lizard", formData.cardFullName);
}

// submit listeners
document
  .getElementById("card-print")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const formData = validateCardForm();
    const zpl = renderZplTemplate(formData);
    const res = await printLabel(zpl);
    await displayPrintNotification(res, "card-result");
  });
document
  .getElementById("card-preview")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const formData = validateCardForm();
    const zpl = renderZplTemplate(formData);
    const res = await requestPreview(zpl);
    await displayPreviewNotification(res, "card-result");
    await displayPreview(res, "card-output");
  });
document
  .getElementById("zpl-print")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const zpl = validateZplForm();
    const res = await printLabel(zpl);
    await displayPrintNotification(res, "zpl-result");
  });
document
  .getElementById("zpl-preview")
  .addEventListener("click", async function (e) {
    e.preventDefault();
    const zpl = validateZplForm();
    const res = await requestPreview(zpl);
    await displayPreviewNotification(res, "zpl-result");
    await displayPreview(res, "zpl-output");
  });
