function getUppercase() {
    let val = document.getElementById("get_key").value;
    let body = {
        s: val
    }

    console.log(JSON.stringify(body))

    fetch("/uppercase", {
        method: 'post',
        body: JSON.stringify(body)
    })
    .then((response) => response.json())
    .then((data) => {
        console.log(data)
        document.getElementById("get_answer").innerHTML = data.v;
    });
}

function getCount() {
    let val = document.getElementById("get_key").value;
    let body = {
        s: val
    }

    console.log(JSON.stringify(body))

    fetch("/count", {
        method: 'post',
        body: JSON.stringify(body)
    })
    .then((response) => response.json())
    .then((data) => {
        console.log(data)
        document.getElementById("get_answer").innerHTML = data.v;
    });
}
