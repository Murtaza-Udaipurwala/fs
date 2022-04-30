const url = document.URL;
const form = document.querySelector("form");
const result = document.getElementById("result");
const btn = document.querySelector("button");
const qrcode = document.getElementById("qrcode");

document.getElementById("url").innerText = url;

const isChecked = () => document.getElementById("onetime").checked;

result.style.visibility = "hidden";

const populate = (data, uploading, err) => {
    if (err) {
        content = "<p>an error occured</p>";
    } else if (uploading) {
        content = "<p>Please be patient. Uploading...</p>";
    } else {
        content = `<p>URL: ${data["url"]}</p>
<p>one time: ${data["onetime"]}</p>
<p>expiry: ${data["expiry"]}</p>
`;

        qrcode.innerHTML = null;
        new QRCode(qrcode, data["url"]);
    }

    result.innerHTML = content;
    result.style.visibility = "visible";
};

const maxSize = 1024 * 1024 * 10;

const validateSize = () => {
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

form.addEventListener("submit", (e) => {
    e.preventDefault();

    btn.disabled = true;

    populate(null, true, false);

    const formData = new FormData();

    if (!validateSize()) {
        form.reset();
        return;
    }

    const file = document.getElementById("file").files[0];
    formData.append("file", file);

    if (isChecked()) {
        formData.append("onetime", "1");
    }

    fetch(`${url}?json=1`, {
        method: "POST",
        body: formData,
    })
        .then((resp) => {
            if (!resp.ok) {
                populate(null, false, true);
                return;
            }

            return resp.json();
        })
        .then((data) => {
            populate(data, false, false);
            btn.disabled = false;
        })
        .catch(() => populate(null, false, true));

    form.reset();
});
