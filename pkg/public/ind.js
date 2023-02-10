let buttonGetImg = document.getElementById("getImg")

const getImg = async (imgID, imgRoute) => {
    let img = {
        "id":imgID,
        "name":imgRoute
    }
    let res = await fetch(`http://localhost:40000/image/getImage`, {
        method:'POST',
        headers:{
           'Content-type':'application/json;charset=utf-8',
        },
        body: JSON.stringify(img)
    })
    if(res.ok){
        let imgWrapper = document.getElementById("images")
        const buffer = await res.arrayBuffer();
        const bytes = new Uint8Array(buffer);
        const blob = new Blob([bytes.buffer]);
        console.log(blob)
        const image = document.createElement('img');
        const reader = new FileReader();
        reader.addEventListener('load', (e) => {
            image.src = e.target.result;
            imgWrapper.appendChild(image)
        });
        reader.readAsDataURL(blob);
    }
}

const getImages = async () => {
    fetch('http://localhost:40000/image/getImages')
        .then((response) => {
            return response.json();
        })
        .then(async (data) => {
            console.log(data);
            for(let i =0; i < data.length; i++){
                await getImg(data[i].id, data[i].name)
            }
        });
}
buttonGetImg.addEventListener("click", getImages)