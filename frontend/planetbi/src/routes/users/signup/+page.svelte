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
      },
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
  <a class="i-have-an-account" href="/users/signin">I have an account</a>

  <div class="form">
    <div class="input-wrapper">
      <input
        type="text"
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

      <svg
        width="16"
        height="11"
        viewBox="0 0 16 11"
        fill="none"
        aria-hidden="true"
        ><path
          d="M14.9421 5.61977L14.4918 5.4024V5.4024L14.9421 5.61977ZM1.05765 5.61977L1.50793 5.4024L1.50793 5.4024L1.05765 5.61977ZM14.9421 5.38039L15.3924 5.16301L15.3924 5.16301L14.9421 5.38039ZM1.05765 5.38039L0.607377 5.16301L0.607377 5.16301L1.05765 5.38039ZM14.4918 5.4024C11.8689 10.8355 4.1308 10.8355 1.50793 5.4024L0.607378 5.83715C3.59408 12.0239 12.4057 12.0239 15.3924 5.83715L14.4918 5.4024ZM14.4918 5.59776C14.462 5.53605 14.462 5.46411 14.4918 5.4024L15.3924 5.83715C15.4952 5.62419 15.4952 5.37596 15.3924 5.16301L14.4918 5.59776ZM1.50793 5.59776C4.1308 0.164679 11.8689 0.164678 14.4918 5.59776L15.3924 5.16301C12.4057 -1.02374 3.59408 -1.02374 0.607377 5.16301L1.50793 5.59776ZM0.607377 5.16301C0.504572 5.37597 0.504572 5.6242 0.607378 5.83715L1.50793 5.4024C1.53772 5.46411 1.53772 5.53605 1.50793 5.59776L0.607377 5.16301ZM9.99987 5.50008C9.99987 6.60465 9.10444 7.50008 7.99987 7.50008V8.50008C9.65672 8.50008 10.9999 7.15693 10.9999 5.50008H9.99987ZM7.99987 7.50008C6.8953 7.50008 5.99987 6.60465 5.99987 5.50008H4.99987C4.99987 7.15693 6.34302 8.50008 7.99987 8.50008V7.50008ZM5.99987 5.50008C5.99987 4.39551 6.8953 3.50008 7.99987 3.50008V2.50008C6.34302 2.50008 4.99987 3.84322 4.99987 5.50008H5.99987ZM7.99987 3.50008C9.10444 3.50008 9.99987 4.39551 9.99987 5.50008H10.9999C10.9999 3.84322 9.65672 2.50008 7.99987 2.50008V3.50008Z"
        ></path></svg
      >
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

      <svg
        width="16"
        height="11"
        viewBox="0 0 16 11"
        fill="none"
        aria-hidden="true"
        ><path
          d="M14.9421 5.61977L14.4918 5.4024V5.4024L14.9421 5.61977ZM1.05765 5.61977L1.50793 5.4024L1.50793 5.4024L1.05765 5.61977ZM14.9421 5.38039L15.3924 5.16301L15.3924 5.16301L14.9421 5.38039ZM1.05765 5.38039L0.607377 5.16301L0.607377 5.16301L1.05765 5.38039ZM14.4918 5.4024C11.8689 10.8355 4.1308 10.8355 1.50793 5.4024L0.607378 5.83715C3.59408 12.0239 12.4057 12.0239 15.3924 5.83715L14.4918 5.4024ZM14.4918 5.59776C14.462 5.53605 14.462 5.46411 14.4918 5.4024L15.3924 5.83715C15.4952 5.62419 15.4952 5.37596 15.3924 5.16301L14.4918 5.59776ZM1.50793 5.59776C4.1308 0.164679 11.8689 0.164678 14.4918 5.59776L15.3924 5.16301C12.4057 -1.02374 3.59408 -1.02374 0.607377 5.16301L1.50793 5.59776ZM0.607377 5.16301C0.504572 5.37597 0.504572 5.6242 0.607378 5.83715L1.50793 5.4024C1.53772 5.46411 1.53772 5.53605 1.50793 5.59776L0.607377 5.16301ZM9.99987 5.50008C9.99987 6.60465 9.10444 7.50008 7.99987 7.50008V8.50008C9.65672 8.50008 10.9999 7.15693 10.9999 5.50008H9.99987ZM7.99987 7.50008C6.8953 7.50008 5.99987 6.60465 5.99987 5.50008H4.99987C4.99987 7.15693 6.34302 8.50008 7.99987 8.50008V7.50008ZM5.99987 5.50008C5.99987 4.39551 6.8953 3.50008 7.99987 3.50008V2.50008C6.34302 2.50008 4.99987 3.84322 4.99987 5.50008H5.99987ZM7.99987 3.50008C9.10444 3.50008 9.99987 4.39551 9.99987 5.50008H10.9999C10.9999 3.84322 9.65672 2.50008 7.99987 2.50008V3.50008Z"
        ></path></svg
      >
    </div>

    <div class="agree-to-terms">
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
