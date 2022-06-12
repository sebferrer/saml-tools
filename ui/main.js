window.onload = function () {
    document.getElementById("decryptButton").onclick = function () {
        let payload = {
            saml_response: document.getElementById("xml").value,
            private_key: document.getElementById("key").value
        };

        let apiURL = '/api/decrypt';

        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function () {
            if (this.readyState == 4 && this.status == 200) {
                document.getElementById("decryptedXML").value = this.responseText;
            }
        };

        xhttp.open("POST", apiURL, true);
        xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhttp.send(JSON.stringify(payload));
    }
};