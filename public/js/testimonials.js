const testimonialData = [
  new Testimonial(
    "60.000.000",
    "laptop rog 4x 23",
    "https://thumb.viva.co.id/media/frontend/thumbs3/2020/05/10/5eb6f370ee3b5-laptop-gaming-msi_1265_711.jpeg",
    "harga ",
    4
  ),
  new Testimonial(
    "60.000.000",
    "laptop rog 4x 23",
    "https://cdn.antaranews.com/cache/730x487/2023/03/05/00.jpg",
    "harga ",
    5
  ),
  new Testimonial(
    "20.0000",
    "laptop asusz gaming",
    "https://idgameware.com/wp-content/uploads/2019/07/201901AM260000004_15487116613591050048825-1.jpg.webp",
    "harga",
    2
  ),
  new Testimonial(
    "60.000.000",
    "laptop rog 4x 23",
    "https://www.asus.com/media/Odin/Websites/global/ProductLine/20200824120814.jpg",
    "harga ",
    1
  ),
  new Testimonial(
    "60.000.000",
    "laptop rog 4x 23",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRvEplX3HYz_fG8S2f1BTgPGaIA4BYr1FHytQ&usqp=CAU",
    "harga ",
    3
  ),
];

// Fungsi untuk merender testimonial ke dalam HTML
function renderTestimonials(testimonials) {
  const testimonialsContainer = document.querySelector(".testimonials");
  testimonialsContainer.innerHTML = ""; // Menghapus testimonial sebelumnya

  if (testimonials.length === 0) {
    document.getElementById("no-testimonial").style.display = "block";
  } else {
    document.getElementById("no-testimonial").style.display = "none";
    testimonials.forEach((testimonial) => {
      testimonialsContainer.innerHTML += testimonial.testimonialHTML;
    });
  }
}

// Fungsi untuk menyaring testimonial berdasarkan rating
function filterTestimonialsByRating(rating) {
  const filteredTestimonials = testimonialData.filter((testimonial) => testimonial.rating === rating);
  renderTestimonials(filteredTestimonials);
}

// Fungsi untuk menampilkan semua testimonial
function showAllTestimonials() {
  renderTestimonials(testimonialData);
}

// Menambahkan event listener pada tombol rating
const ratingButtons = document.querySelectorAll(".rating-button");
ratingButtons.forEach((button) => {
  button.addEventListener("click", () => {
    const rating = parseInt(button.dataset.rating);
    filterTestimonialsByRating(rating);
  });
});

// Menambahkan event listener pada tombol "All Ratings"
const allRatingsButton = document.querySelector(".all-ratings-button");
allRatingsButton.addEventListener("click", showAllTestimonials);

// Merender testimonial secara awal
renderTestimonials(testimonialData);