const formData = new FormData();

formData.append("title", "Counting");
formData.append("description", "Learn How To Count");
formData.append("price", "2000");
formData.append("language", "en");
formData.append("level", "bigener");
formData.append("duration", "5");
formData.append("video", fs.createReadStream("./assets/123.mp4"));
formData.append("image", fs.createReadStream("./assets/123.jpg"));

const response = await fetchD(
	ur
);

console.log(response.data);