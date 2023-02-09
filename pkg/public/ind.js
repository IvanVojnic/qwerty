let buttonGetImg = document.getElementById("getImg")

const getImg = async (imgID) => {
    let res = await fetch(`http://localhost:40000/image/getImages`, {
        method:'POST',
        headers:{
           'Content-type':'application/json;charset=utf-8',
        },
        body: JSON.stringify(imgID)
    })
    if(res.ok){
        let imgWrapper = document.getElementById("images")
        let currentImgBox = res.body
        let currentImg = currentImgBox.imgBody
        const buffer = await currentImg.arrayBuffer();
        const bytes = new Uint8Array(buffer);
        const blob = new Blob([bytes.buffer]);

        const image = document.createElement('img');
        const reader = new FileReader();

        reader.addEventListener('load', (e) => {
            image.src = e.target.result;
            this.$el.append(image);
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
                await getImg(data[i].id)
            }
        });
}
buttonGetImg.addEventListener("click", getImages)