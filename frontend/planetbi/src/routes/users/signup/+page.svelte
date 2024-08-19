<script>
  import { title } from "$lib/stores.js";
  title.set("Signup");

  import PlanetBIDarkSVG from "$lib/assets/planetbi-dark.svg";

  import { hasValue } from "$lib/common.js";
  import { http } from "$lib/http.js";

  let name;
  let email;
  let password;
  let confirmPassword;
  let error = "";
  let success = false;
  let disabled = false;
  let running = false;

  function submitOnEnter(e) {
    if (e.keyCode === 13) {
      signup();
    }
  }

  // https://stackoverflow.com/a/13975255
  function isValidEmail(value) {
    let input = document.createElement("input");
    input.type = "email";
    input.value = value;
    if (typeof input.checkValidity === "function") {
      if (!input.checkValidity()) {
        return false;
      }
    }
    return /\S+@\S+\.\S+/.test(value);
  }

  function signup() {
    error = "";
    disabled = true;
    running = true;

    // validate name
    if (!hasValue(name)) {
      error = "Please enter your name.";
      disabled = false;
      running = false;
      return;
    }

    // validate email
    if (!hasValue(email) || !isValidEmail(email)) {
      error = "Please enter a valid email.";
      disabled = false;
      running = false;
      return;
    }

    // validate password
    // don't add complexity requirements
    // https://learn.microsoft.com/en-us/windows/security/threat-protection/security-policy-settings/password-must-meet-complexity-requirements
    if (!hasValue(password) || password.length < 8) {
      error = "Password must be 8+ chars.";
      disabled = false;
      running = false;
      return;
    }

    // confirm password
    if (password !== confirmPassword) {
      error = "Passwords do not match.";
      disabled = false;
      running = false;
      return;
    }

    // clear error
    error = "";

    // signup
    http.post(
      "/signup",
      { name, email, password },
      function (res, status) {
        success = true;
      },
      function (res, status) {
        switch (status) {
          case 409:
            error = "Email address already in use.";
            break;
          case 408:
            error = "Invalid email.";
            break;
          case 407:
            // we don't have to specify the error here
            error = "Invalid password.";
            break;
          default:
            error = "Unexpected error, please try again.";
            break;
        }

        disabled = false;
        running = false;
      }
    );
  }
</script>

<div class="account-created {!success ? 'hide' : ''}">
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="50"
    height="50"
    enable-background="new 0 0 256 256"
    viewBox="0 0 256 256"
  >
    <linearGradient
      id="aa"
      x1="127.999"
      x2="127.999"
      y1="67.564"
      y2="224.484"
      gradientUnits="userSpaceOnUse"
    >
      <stop offset="0" stop-color="#5ae050" />
      <stop offset="1" stop-color="#40d437" />
    </linearGradient>
    <path
      fill="url(#aa)"
      fill-rule="evenodd"
      d="M27.364,99.627c3.874-4.29,10.742-4.86,15.268-1.26   l52.917,42.091c4.526,3.601,11.638,3.27,15.81-0.732L213.38,41.82c4.171-4.002,10.887-3.891,14.92,0.254l24.74,25.404   c4.036,4.139,3.931,10.809-0.236,14.814L116.078,213.814c-4.166,4.008-11.336,4.417-15.932,0.911L3.873,141.291   c-4.595-3.507-5.188-9.883-1.313-14.179L27.364,99.627z"
      clip-rule="evenodd"
    />
    <path
      fill="#2C2C2C"
      fill-rule="evenodd"
      d="M252.804,77.914L116.078,209.437    c-4.166,4.008-11.336,4.417-15.932,0.91L3.873,136.912c-1.903-1.452-3.107-3.397-3.604-5.483    c-0.854,3.563,0.395,7.413,3.604,9.862l96.273,73.435c4.596,3.506,11.766,3.097,15.932-0.911L252.804,82.292    c2.69-2.586,3.679-6.283,2.959-9.692C255.342,74.562,254.355,76.423,252.804,77.914z"
      clip-rule="evenodd"
      opacity=".2"
    />
  </svg>
  <h1>Account created</h1>
  <p>Please check your email to verify your account.</p>
</div>

<div class="header">
  <a href="/" style="text-decoration: none;" class="logo">
    <img src={PlanetBIDarkSVG} alt="PlanetBI" />
  </a>

  <div>Planet<span>BI</span></div>
</div>

<div class="form-wrapper {success ? 'hide' : ''}">
  <div class="title">Sign up</div>
  <a class="ihaveanaccount" href="/users/signin">I have an account</a>

  <div class="form">
    <div class="input-wrapper">
      <input
        type="name"
        class={hasValue(name) ? "has-value" : ""}
        on:keydown={submitOnEnter}
        bind:value={name}
        {disabled}
      />
      <span class="placeholder">Name</span>
    </div>
    <div class="input-wrapper">
      <input
        type="email"
        class={hasValue(email) ? "has-value" : ""}
        on:keydown={submitOnEnter}
        bind:value={email}
        {disabled}
      />
      <span class="placeholder">Email</span>
    </div>
    <div class="input-wrapper">
      <input
        type="password"
        class={hasValue(password) ? "has-value" : ""}
        on:keydown={submitOnEnter}
        bind:value={password}
        {disabled}
      />
      <span class="placeholder">Password</span>
    </div>
    <div class="input-wrapper">
      <input
        type="password"
        class={hasValue(confirmPassword) ? "has-value" : ""}
        on:keydown={submitOnEnter}
        bind:value={confirmPassword}
        {disabled}
      />
      <span class="placeholder">Confirm password</span>
    </div>

    <div class="agreetoterms">
      By signing up, you agree to the <a href="/terms-and-conditions"
        >Terms and Conditions</a
      >.
    </div>

    <button
      class="signup btn-blue ld-over {running ? 'running' : ''}"
      id="signup"
      on:click={signup}
      {disabled}
    >
      <div class="ld ld-ring ld-cycle"></div>
      Agree and Sign up
    </button>

    <div class="error {!error ? 'hide' : ''}">
      <span>{error}</span>
      <button
        on:click={() => {
          error = "";
        }}>Okay</button
      >
    </div>
  </div>
</div>

<style>
  @import "signup.css";
</style>
