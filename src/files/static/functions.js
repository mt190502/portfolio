const light_theme_logo =
  "M6.995 12c0 2.761 2.246 5.007 5.007 5.007s5.007-2.246 5.007-5.007-2.246-5.007-5.007-5.007S6.995 9.239 6.995 12zM11 19h2v3h-2zm0-17h2v3h-2zm-9 9h3v2H2zm17 0h3v2h-3zM5.637 19.778l-1.414-1.414 2.121-2.121 1.414 1.414zM16.242 6.344l2.122-2.122 1.414 1.414-2.122 2.122zM6.344 7.759 4.223 5.637l1.415-1.414 2.12 2.122zm13.434 10.605-1.414 1.414-2.122-2.122 1.414-1.414z";
const dark_theme_logo =
  "M12 11.807A9.002 9.002 0 0 1 10.049 2a9.942 9.942 0 0 0-5.12 2.735c-3.905 3.905-3.905 10.237 0 14.142 3.906 3.906 10.237 3.905 14.143 0a9.946 9.946 0 0 0 2.735-5.119A9.003 9.003 0 0 1 12 11.807z";

function switchTheme() {
  let themeIcon = "";
  var currentTheme = document.documentElement.getAttribute("data-theme");
  if (currentTheme === "light") {
    document.documentElement.setAttribute("data-theme", "dark");
    localStorage.setItem("theme", "dark");
    themeIcon = dark_theme_logo;
  }
  if (currentTheme === "dark") {
    document.documentElement.setAttribute("data-theme", "light");
    localStorage.setItem("theme", "light");
    themeIcon = light_theme_logo;
  }
  const btn = document.querySelector("#toggle_theme");
  btn?.setAttribute("d", themeIcon);
}

function checkEmail(emailID) {
  const email = document.querySelector(emailID);

  return {
    email: email?.value,
    valid: email?.value.match(/^[^\s@]+@[^\s@]+\.[^\s@]+$/) ? true : false,
    length: email?.value.length,
    path: email,
  };
}

function checkPassword(passwordID) {
  const password = document.querySelector(passwordID);

  return {
    password: password?.value,
    length: password?.value.length,
    path: password,
  };
}

function checkLoginCred() {
  const email = checkEmail("#email");
  const password = checkPassword("#password");
  const submit = document.querySelector("#submit");
  const submitEnable = email.valid && password.length ? true : false;

  if (!email.valid) {
    email.path.classList.add("input-error");
  } else {
    email.path.classList.remove("input-error");
  }

  if (submitEnable) {
    submit?.removeAttribute("disabled");
  } else {
    submit?.setAttribute("disabled", "true");
  }
}

function checkRegisterCred() {
  const email = checkEmail("#email");
  const password = checkPassword("#password");
  const confirm = checkPassword("#confirm");
  const submit = document.querySelector("#submit");
  const submitEnable =
    email.valid && password.length > 0 && password.password == confirm.password
      ? true
      : false;

  if (!email.valid) {
    email.path.classList.add("input-error");
  } else {
    email.path.classList.remove("input-error");
  }

  if (
    password.length != 0 &&
    confirm.length != 0 &&
    password.password != confirm.password
  ) {
    password.path.classList.add("input-error");
    confirm.path.classList.add("input-error");
  } else {
    password.path.classList.remove("input-error");
    confirm.path.classList.remove("input-error");
  }

  if (submitEnable) {
    submit?.removeAttribute("disabled");
  } else {
    submit?.setAttribute("disabled", "true");
  }
}

function checkNewEmail() {
  const email = checkEmail("#new-email");
  const confirm = checkEmail("#confirm-email");
  const submit = document.querySelector("#new-email-submit");
  const submitEnable =
    email.valid && confirm.valid && email.email == confirm.email ? true : false;

  if (!email.valid && email.length != 0) {
    email.path.classList.remove("input-solid");
    email.path.classList.add("input-error");
  } else {
    email.path.classList.remove("input-error");
    email.path.classList.add("input-solid");
  }

  if (!confirm.valid && confirm.length != 0) {
    confirm.path.classList.remove("input-solid");
    confirm.path.classList.add("input-error");
  } else {
    confirm.path.classList.remove("input-error");
    confirm.path.classList.add("input-solid");
  }

  if (email.valid && confirm.valid && email.email != confirm.email) {
    email.path.classList.remove("input-solid");
    confirm.path.classList.remove("input-solid");
    email.path.classList.add("input-error");
    confirm.path.classList.add("input-error");
  }
  if (email.valid && confirm.valid && email.email == confirm.email) {
    email.path.classList.remove("input-error");
    confirm.path.classList.remove("input-error");
    email.path.classList.add("input-solid");
    confirm.path.classList.add("input-solid");
  }

  if (submitEnable) {
    submit?.removeAttribute("disabled");
  } else {
    submit?.setAttribute("disabled", "true");
  }
}

function checkNewPassword() {
  const password = checkPassword("#new-password");
  const confirm = checkPassword("#confirm-password");
  const submit = document.querySelector("#new-password-submit");
  const submitEnable =
    password.length && password.password == confirm.password ? true : false;

  if (
    password.length != 0 &&
    confirm.length != 0 &&
    password.password != confirm.password
  ) {
    password.path.classList.remove("input-solid");
    confirm.path.classList.remove("input-solid");
    password.path.classList.add("input-error");
    confirm.path.classList.add("input-error");
  } else {
    password.path.classList.remove("input-error");
    confirm.path.classList.remove("input-error");
    password.path.classList.add("input-solid");
    confirm.path.classList.add("input-solid");
  }

  if (submitEnable) {
    submit?.removeAttribute("disabled");
  } else {
    submit?.setAttribute("disabled", "true");
  }
}

(() => {
  document.getElementById("error-notification")?.classList.add("slide-out");
  document.getElementById("success-notification")?.classList.add("slide-out");

  const storedTheme = localStorage.getItem("theme");
  if (storedTheme === null) {
    localStorage.setItem("theme", "dark");
    document.documentElement.setAttribute("data-theme", "dark");
  } else {
    document.documentElement.setAttribute("data-theme", storedTheme);
  }

  const btn = document.querySelector("#toggle_theme");
  if (localStorage.getItem("theme") === "dark") {
    btn?.setAttribute("d", dark_theme_logo);
  } else {
    btn?.setAttribute("d", light_theme_logo);
  }
})();
