const formElement = document.getElementById('custom-form');
console.log(formElement)
const formAction = async (e) => {
  e.preventDefault();
  const form = e.target;
  const endpoint = form.action;
  const formData = new FormData(form);

  const user = {};
  formData.forEach((value, key) => {
    user[key] = value;
  });
  console.log(user)

  try {
    const response = await fetch(endpoint, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json; charset=UTF-8'
      },
      body: JSON.stringify(user)
    });

    const jsonResponse = await response.json();
    console.log(jsonResponse);
    if (jsonResponse.success) {
      alert("Your message is well received! 👋!")
      window.location.href = 'https://show-resume.herokuapp.com';
    }
    console.log(jsonResponse.message);
  } catch (error) {
    console.log(error);
  }
}

formElement.addEventListener('submit', formAction)
