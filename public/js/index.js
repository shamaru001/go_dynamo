fetch('http://localhost:8000/getprofile')
  .then(response => response.json())
  .then(data => {

    const Name = document.querySelector("#Name");
    const Description = document.querySelector("#Description");
    const Image = document.querySelector("#Image");

    
    Name.innerHTML = data.Name;
    Description.innerHTML = data.Description;
    Image.src = data.Image;

});