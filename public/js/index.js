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


const form = document.querySelector("#updateProfile");

form.addEventListener("submit", async (e) => {
    e.preventDefault();
    
    const Name = document.querySelector("#name").value;
    const Image = document.querySelector("#image").value;
    const Description = document.querySelector("#description").value;


    const rawResponse = await fetch('/updateprofile/1', {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
			Name,
			Image,
			Description,
		})
      });
      const content = await rawResponse.json();

	  window.location.reload(true)
});
