const url = document.URL;
const form = document.querySelector("form");
const result = document.getElementById("result");
const btn = document.querySelector("button");
const qrcode = document.getElementById("qrcode");
const maxSize = 1024 * 1024 * 10;

// set referenced url
document.querySelector(".url").innerText = url;

// hide result region
result.style.visibility = "hidden";

function isChecked() {
    document.getElementById("onetime").checked;
}

function resetQRCode() {
    qrcode.innerHTML = null;
}

function generateQRCode(url) {
    new QRCode(qrcode, url);
}

function feedback(text) {
    result.innerHTML = `<p>${text}</p>`;
    result.style.visibility = "visible";
}

function getFormattedResult(data) {
    content = `<p>URL: ${data["url"]}</p>
<p>one time: ${data["onetime"]}</p>
<p>expiry: ${data["expiry"]}</p>
`;

    return content;
}

function validateSize() {
    const ifile = document.getElementById("file");

    if (ifile.files.length === 0) {
        return false;
    }

    const size = ifile.files[0].size;

    if (size > maxSize) {
        alert("File too big, please select a file less than 10MiB");
        return false;
    }

    return true;
};

function toggleForm(disable) {
    const inputs = document.querySelectorAll("form input");
    inputs.forEach(input => {
        input.disabled = disable;
    });

    const buttons = document.querySelectorAll("form button");
    buttons.forEach(btn => {
        btn.disabled = disable;
    });
}

function resetForm() {
    form.reset();
}

async function upload(formData) {
    const resp = await fetch(`${url}?json=1`, {
        method: "POST",
        body: formData
    });

    if (!resp.ok) {
        feedback("An error occured");
        return;
    }

    const data = await resp.json();

    const content = getFormattedResult(data);
    feedback(content);

    return data;
}

form.addEventListener("submit", (e) => {
    e.preventDefault();
    toggleForm(true);

    const formData = new FormData();

    if (!validateSize()) {
        toggleForm(false);
        resetForm();
        return;
    }

    const file = document.getElementById("file").files[0];
    formData.append("file", file);

    if (isChecked()) {
        formData.append("onetime", "1");
    }

    resetQRCode();
    feedback("Please be patient. Uploading...");

    upload(formData).then((data) => {
        if (!data) return;

        generateQRCode(data["url"]);
        toggleForm(false);
        resetForm();
    });
});