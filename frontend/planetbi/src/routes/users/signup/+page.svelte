<script>
  import { hasValue } from "$lib/js/common.js";
  import { http } from "$lib/js/http.js";
  import { page } from "$app/stores";
  import Header from "$lib/components/auth-header/component.svelte";
  import { title } from "$lib/js/stores.js";
  title.set("Sign up");

  let name;
  let email;
  let password;
  let confirmPassword;
  let error = "";
  let success = false;
  let disabled = false;
  let running = false;
  let revealPassword = false;

  function onRouteChange() {
    if ($page.url.hash == "#success") {
      success = true;
    } else {
      success = false;
    }
  }

  function submitOnEnter(e) {
    if (e.key === "Enter") {
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
    if (!hasValue(email)) {
      error = "Please enter your email.";
      disabled = false;
      running = false;
      return;
    }
    if (!isValidEmail(email)) {
      error = "Please enter a valid email.";
      disabled = false;
      running = false;
      return;
    }

    // validate password
    // don't add complexity requirements
    // https://learn.microsoft.com/en-us/windows/security/threat-protection/security-policy-settings/password-must-meet-complexity-requirements
    if (!hasValue(password)) {
      error = "Please enter your password.";
      disabled = false;
      running = false;
      return;
    }
    if (password.length < 8) {
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
        disabled = false;
        running = false;
        name = "";
        email = "";
        password = "";
        confirmPassword = "";
        revealPassword = false;
      },
      function (res, status) {
        switch (status) {
          case 408:
            error = "Invalid email address.";
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

<svelte:window on:hashchange={onRouteChange} />

<Header />

<div class="form-wrapper">
  <div class="account-created {!success ? 'hide' : ''}">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      xmlns:xlink="http://www.w3.org/1999/xlink"
      version="1.1"
      viewBox="0 0 512 512"
      xml:space="preserve"
    >
      <g>
        <path
          d="M510.746,110.361c-2.128-10.754-6.926-20.918-13.926-29.463c-1.422-1.794-2.909-3.39-4.535-5.009   c-12.454-12.52-29.778-19.701-47.531-19.701H67.244c-17.951,0-34.834,7-47.539,19.708c-1.608,1.604-3.099,3.216-4.575,5.067   c-6.97,8.509-11.747,18.659-13.824,29.428C0.438,114.62,0,119.002,0,123.435v265.137c0,9.224,1.874,18.206,5.589,26.745   c3.215,7.583,8.093,14.772,14.112,20.788c1.516,1.509,3.022,2.901,4.63,4.258c12.034,9.966,27.272,15.45,42.913,15.45h377.51   c15.742,0,30.965-5.505,42.967-15.56c1.604-1.298,3.091-2.661,4.578-4.148c5.818-5.812,10.442-12.49,13.766-19.854l0.438-1.05   c3.646-8.377,5.497-17.33,5.497-26.628V123.435C512,119.06,511.578,114.649,510.746,110.361z M34.823,99.104   c0.951-1.392,2.165-2.821,3.714-4.382c7.689-7.685,17.886-11.914,28.706-11.914h377.51c10.915,0,21.115,4.236,28.719,11.929   c1.313,1.327,2.567,2.8,3.661,4.272l2.887,3.88l-201.5,175.616c-6.212,5.446-14.21,8.443-22.523,8.443   c-8.231,0-16.222-2.99-22.508-8.436L32.19,102.939L34.823,99.104z M26.755,390.913c-0.109-0.722-0.134-1.524-0.134-2.341V128.925   l156.37,136.411L28.199,400.297L26.755,390.913z M464.899,423.84c-6.052,3.492-13.022,5.344-20.145,5.344H67.244   c-7.127,0-14.094-1.852-20.142-5.344l-6.328-3.668l159.936-139.379l17.528,15.246c10.514,9.128,23.922,14.16,37.761,14.16   c13.89,0,27.32-5.032,37.827-14.16l17.521-15.253L471.228,420.18L464.899,423.84z M485.372,388.572   c0,0.803-0.015,1.597-0.116,2.304l-1.386,9.472L329.012,265.409l156.36-136.418V388.572z"
        />
      </g>
    </svg>

    <h1>Almost done!</h1>
    <p>Please check your email to finish the process.</p>
  </div>

  <div class="title {success ? 'hide' : ''}">Sign up</div>
  <a class="i-have-an-account {success ? 'hide' : ''}" href="/users/signin"
    >I have an account</a
  >

  <div class="form {success ? 'hide' : ''}">
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
      {#if revealPassword}
        <input
          type="text"
          class={hasValue(password) ? "has-value" : ""}
          on:keydown={submitOnEnter}
          bind:value={password}
          {disabled}
        />
        <span class="placeholder">Password</span>

        <svg
          viewBox="0 0 16 14"
          fill="none"
          aria-hidden="true"
          class="revealed"
          on:click={() => (revealPassword = false)}
          ><path
            d="M14.9421 7.6197L14.4918 7.40232V7.40232L14.9421 7.6197ZM1.05765 7.6197L1.50793 7.40232L1.50793 7.40232L1.05765 7.6197ZM14.9421 7.38031L15.3924 7.16294L15.3924 7.16293L14.9421 7.38031ZM1.05765 7.38031L0.607377 7.16293L0.607377 7.16294L1.05765 7.38031ZM10.9999 7.5C10.9999 7.22386 10.776 7 10.4999 7C10.2237 7 9.99987 7.22386 9.99987 7.5H10.9999ZM7.99987 9.5C7.72373 9.5 7.49987 9.72386 7.49987 10C7.49987 10.2761 7.72373 10.5 7.99987 10.5V9.5ZM1.14632 13.1464C0.951062 13.3417 0.951062 13.6583 1.14632 13.8536C1.34159 14.0488 1.65817 14.0488 1.85343 13.8536L1.14632 13.1464ZM14.8534 0.853553C15.0487 0.658291 15.0487 0.341709 14.8534 0.146447C14.6582 -0.0488155 14.3416 -0.0488155 14.1463 0.146447L14.8534 0.853553ZM11.0879 4.20226C11.3382 4.31901 11.6357 4.21079 11.7524 3.96055C11.8692 3.7103 11.761 3.41279 11.5107 3.29603L11.0879 4.20226ZM13.3114 4.45031C13.1009 4.2715 12.7854 4.29713 12.6066 4.50756C12.4278 4.71799 12.4534 5.03353 12.6638 5.21234L13.3114 4.45031ZM3.78294 11.331C4.0204 11.4719 4.32716 11.3937 4.46811 11.1562C4.60906 10.9187 4.53082 10.612 4.29335 10.471L3.78294 11.331ZM6.12685 11.2361C5.85976 11.166 5.58639 11.3257 5.51627 11.5928C5.44614 11.8599 5.60582 12.1332 5.87291 12.2033L6.12685 11.2361ZM5.5293 9.20225C5.68614 9.42953 5.99753 9.48663 6.22481 9.32978C6.45209 9.17294 6.50918 8.86155 6.35234 8.63427L5.5293 9.20225ZM9.09968 5.82922C9.33016 5.98132 9.6403 5.91778 9.7924 5.6873C9.9445 5.45682 9.88096 5.14668 9.65048 4.99458L9.09968 5.82922ZM14.4918 7.59768C14.462 7.53597 14.462 7.46404 14.4918 7.40232L15.3924 7.83707C15.4952 7.62411 15.4952 7.37589 15.3924 7.16294L14.4918 7.59768ZM0.607377 7.16294C0.504572 7.37589 0.504572 7.62412 0.607378 7.83707L1.50793 7.40232C1.53772 7.46403 1.53772 7.53597 1.50793 7.59768L0.607377 7.16294ZM9.99987 7.5C9.99987 8.60457 9.10444 9.5 7.99987 9.5V10.5C9.65672 10.5 10.9999 9.15686 10.9999 7.5H9.99987ZM5.99987 7.5C5.99987 6.39543 6.8953 5.5 7.99987 5.5V4.5C6.34302 4.5 4.99987 5.84315 4.99987 7.5H5.99987ZM1.85343 13.8536L14.8534 0.853553L14.1463 0.146447L1.14632 13.1464L1.85343 13.8536ZM1.50793 7.59768C3.35244 3.77692 7.73141 2.63626 11.0879 4.20226L11.5107 3.29603C7.70772 1.52172 2.71273 2.80184 0.607377 7.16293L1.50793 7.59768ZM12.6638 5.21234C13.3942 5.83295 14.0231 6.62685 14.4918 7.59768L15.3924 7.16293C14.8632 6.06675 14.1481 5.16133 13.3114 4.45031L12.6638 5.21234ZM4.29335 10.471C3.15305 9.79421 2.17002 8.7738 1.50793 7.40232L0.607378 7.83707C1.35622 9.38825 2.47624 10.5554 3.78294 11.331L4.29335 10.471ZM14.4918 7.40232C12.8631 10.7762 9.26241 12.0594 6.12685 11.2361L5.87291 12.2033C9.42276 13.1354 13.5347 11.6852 15.3924 7.83707L14.4918 7.40232ZM6.35234 8.63427C6.13001 8.31211 5.99987 7.92201 5.99987 7.5H4.99987C4.99987 8.13132 5.19543 8.71846 5.5293 9.20225L6.35234 8.63427ZM7.99987 5.5C8.40703 5.5 8.78437 5.62114 9.09968 5.82922L9.65048 4.99458C9.17675 4.68196 8.6088 4.5 7.99987 4.5V5.5Z"
            fill="#6D758D"
          ></path></svg
        >
      {:else}
        <input
          type="password"
          class={hasValue(password) ? "has-value" : ""}
          on:keydown={submitOnEnter}
          bind:value={password}
          {disabled}
        />
        <span class="placeholder">Password</span>

        <svg
          viewBox="0 0 16 11"
          fill="none"
          aria-hidden="true"
          on:click={() => (revealPassword = true)}
          ><path
            d="M14.9421 5.61977L14.4918 5.4024V5.4024L14.9421 5.61977ZM1.05765 5.61977L1.50793 5.4024L1.50793 5.4024L1.05765 5.61977ZM14.9421 5.38039L15.3924 5.16301L15.3924 5.16301L14.9421 5.38039ZM1.05765 5.38039L0.607377 5.16301L0.607377 5.16301L1.05765 5.38039ZM14.4918 5.4024C11.8689 10.8355 4.1308 10.8355 1.50793 5.4024L0.607378 5.83715C3.59408 12.0239 12.4057 12.0239 15.3924 5.83715L14.4918 5.4024ZM14.4918 5.59776C14.462 5.53605 14.462 5.46411 14.4918 5.4024L15.3924 5.83715C15.4952 5.62419 15.4952 5.37596 15.3924 5.16301L14.4918 5.59776ZM1.50793 5.59776C4.1308 0.164679 11.8689 0.164678 14.4918 5.59776L15.3924 5.16301C12.4057 -1.02374 3.59408 -1.02374 0.607377 5.16301L1.50793 5.59776ZM0.607377 5.16301C0.504572 5.37597 0.504572 5.6242 0.607378 5.83715L1.50793 5.4024C1.53772 5.46411 1.53772 5.53605 1.50793 5.59776L0.607377 5.16301ZM9.99987 5.50008C9.99987 6.60465 9.10444 7.50008 7.99987 7.50008V8.50008C9.65672 8.50008 10.9999 7.15693 10.9999 5.50008H9.99987ZM7.99987 7.50008C6.8953 7.50008 5.99987 6.60465 5.99987 5.50008H4.99987C4.99987 7.15693 6.34302 8.50008 7.99987 8.50008V7.50008ZM5.99987 5.50008C5.99987 4.39551 6.8953 3.50008 7.99987 3.50008V2.50008C6.34302 2.50008 4.99987 3.84322 4.99987 5.50008H5.99987ZM7.99987 3.50008C9.10444 3.50008 9.99987 4.39551 9.99987 5.50008H10.9999C10.9999 3.84322 9.65672 2.50008 7.99987 2.50008V3.50008Z"
          ></path></svg
        >
      {/if}
    </div>
    <div class="input-wrapper">
      {#if revealPassword}
        <input
          type="text"
          class={hasValue(confirmPassword) ? "has-value" : ""}
          on:keydown={submitOnEnter}
          bind:value={confirmPassword}
          {disabled}
        />
        <span class="placeholder">Confirm password</span>

        <svg
          viewBox="0 0 16 14"
          fill="none"
          aria-hidden="true"
          class="revealed"
          on:click={() => (revealPassword = false)}
          ><path
            d="M14.9421 7.6197L14.4918 7.40232V7.40232L14.9421 7.6197ZM1.05765 7.6197L1.50793 7.40232L1.50793 7.40232L1.05765 7.6197ZM14.9421 7.38031L15.3924 7.16294L15.3924 7.16293L14.9421 7.38031ZM1.05765 7.38031L0.607377 7.16293L0.607377 7.16294L1.05765 7.38031ZM10.9999 7.5C10.9999 7.22386 10.776 7 10.4999 7C10.2237 7 9.99987 7.22386 9.99987 7.5H10.9999ZM7.99987 9.5C7.72373 9.5 7.49987 9.72386 7.49987 10C7.49987 10.2761 7.72373 10.5 7.99987 10.5V9.5ZM1.14632 13.1464C0.951062 13.3417 0.951062 13.6583 1.14632 13.8536C1.34159 14.0488 1.65817 14.0488 1.85343 13.8536L1.14632 13.1464ZM14.8534 0.853553C15.0487 0.658291 15.0487 0.341709 14.8534 0.146447C14.6582 -0.0488155 14.3416 -0.0488155 14.1463 0.146447L14.8534 0.853553ZM11.0879 4.20226C11.3382 4.31901 11.6357 4.21079 11.7524 3.96055C11.8692 3.7103 11.761 3.41279 11.5107 3.29603L11.0879 4.20226ZM13.3114 4.45031C13.1009 4.2715 12.7854 4.29713 12.6066 4.50756C12.4278 4.71799 12.4534 5.03353 12.6638 5.21234L13.3114 4.45031ZM3.78294 11.331C4.0204 11.4719 4.32716 11.3937 4.46811 11.1562C4.60906 10.9187 4.53082 10.612 4.29335 10.471L3.78294 11.331ZM6.12685 11.2361C5.85976 11.166 5.58639 11.3257 5.51627 11.5928C5.44614 11.8599 5.60582 12.1332 5.87291 12.2033L6.12685 11.2361ZM5.5293 9.20225C5.68614 9.42953 5.99753 9.48663 6.22481 9.32978C6.45209 9.17294 6.50918 8.86155 6.35234 8.63427L5.5293 9.20225ZM9.09968 5.82922C9.33016 5.98132 9.6403 5.91778 9.7924 5.6873C9.9445 5.45682 9.88096 5.14668 9.65048 4.99458L9.09968 5.82922ZM14.4918 7.59768C14.462 7.53597 14.462 7.46404 14.4918 7.40232L15.3924 7.83707C15.4952 7.62411 15.4952 7.37589 15.3924 7.16294L14.4918 7.59768ZM0.607377 7.16294C0.504572 7.37589 0.504572 7.62412 0.607378 7.83707L1.50793 7.40232C1.53772 7.46403 1.53772 7.53597 1.50793 7.59768L0.607377 7.16294ZM9.99987 7.5C9.99987 8.60457 9.10444 9.5 7.99987 9.5V10.5C9.65672 10.5 10.9999 9.15686 10.9999 7.5H9.99987ZM5.99987 7.5C5.99987 6.39543 6.8953 5.5 7.99987 5.5V4.5C6.34302 4.5 4.99987 5.84315 4.99987 7.5H5.99987ZM1.85343 13.8536L14.8534 0.853553L14.1463 0.146447L1.14632 13.1464L1.85343 13.8536ZM1.50793 7.59768C3.35244 3.77692 7.73141 2.63626 11.0879 4.20226L11.5107 3.29603C7.70772 1.52172 2.71273 2.80184 0.607377 7.16293L1.50793 7.59768ZM12.6638 5.21234C13.3942 5.83295 14.0231 6.62685 14.4918 7.59768L15.3924 7.16293C14.8632 6.06675 14.1481 5.16133 13.3114 4.45031L12.6638 5.21234ZM4.29335 10.471C3.15305 9.79421 2.17002 8.7738 1.50793 7.40232L0.607378 7.83707C1.35622 9.38825 2.47624 10.5554 3.78294 11.331L4.29335 10.471ZM14.4918 7.40232C12.8631 10.7762 9.26241 12.0594 6.12685 11.2361L5.87291 12.2033C9.42276 13.1354 13.5347 11.6852 15.3924 7.83707L14.4918 7.40232ZM6.35234 8.63427C6.13001 8.31211 5.99987 7.92201 5.99987 7.5H4.99987C4.99987 8.13132 5.19543 8.71846 5.5293 9.20225L6.35234 8.63427ZM7.99987 5.5C8.40703 5.5 8.78437 5.62114 9.09968 5.82922L9.65048 4.99458C9.17675 4.68196 8.6088 4.5 7.99987 4.5V5.5Z"
            fill="#6D758D"
          ></path></svg
        >
      {:else}
        <input
          type="password"
          class={hasValue(confirmPassword) ? "has-value" : ""}
          on:keydown={submitOnEnter}
          bind:value={confirmPassword}
          {disabled}
        />
        <span class="placeholder">Confirm password</span>

        <svg
          viewBox="0 0 16 11"
          fill="none"
          aria-hidden="true"
          on:click={() => (revealPassword = true)}
          ><path
            d="M14.9421 5.61977L14.4918 5.4024V5.4024L14.9421 5.61977ZM1.05765 5.61977L1.50793 5.4024L1.50793 5.4024L1.05765 5.61977ZM14.9421 5.38039L15.3924 5.16301L15.3924 5.16301L14.9421 5.38039ZM1.05765 5.38039L0.607377 5.16301L0.607377 5.16301L1.05765 5.38039ZM14.4918 5.4024C11.8689 10.8355 4.1308 10.8355 1.50793 5.4024L0.607378 5.83715C3.59408 12.0239 12.4057 12.0239 15.3924 5.83715L14.4918 5.4024ZM14.4918 5.59776C14.462 5.53605 14.462 5.46411 14.4918 5.4024L15.3924 5.83715C15.4952 5.62419 15.4952 5.37596 15.3924 5.16301L14.4918 5.59776ZM1.50793 5.59776C4.1308 0.164679 11.8689 0.164678 14.4918 5.59776L15.3924 5.16301C12.4057 -1.02374 3.59408 -1.02374 0.607377 5.16301L1.50793 5.59776ZM0.607377 5.16301C0.504572 5.37597 0.504572 5.6242 0.607378 5.83715L1.50793 5.4024C1.53772 5.46411 1.53772 5.53605 1.50793 5.59776L0.607377 5.16301ZM9.99987 5.50008C9.99987 6.60465 9.10444 7.50008 7.99987 7.50008V8.50008C9.65672 8.50008 10.9999 7.15693 10.9999 5.50008H9.99987ZM7.99987 7.50008C6.8953 7.50008 5.99987 6.60465 5.99987 5.50008H4.99987C4.99987 7.15693 6.34302 8.50008 7.99987 8.50008V7.50008ZM5.99987 5.50008C5.99987 4.39551 6.8953 3.50008 7.99987 3.50008V2.50008C6.34302 2.50008 4.99987 3.84322 4.99987 5.50008H5.99987ZM7.99987 3.50008C9.10444 3.50008 9.99987 4.39551 9.99987 5.50008H10.9999C10.9999 3.84322 9.65672 2.50008 7.99987 2.50008V3.50008Z"
          ></path></svg
        >
      {/if}
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
        }}>hide</button
      >
    </div>
  </div>
</div>

<svelte:head>
  <style>
    body {
      background: #f2f7fc;
    }

    .form-wrapper {
      margin: 44px auto 45px;
      width: 450px;
      position: relative;
      background: #fff;
      border-radius: 8px;
      padding: 25px 32px 9px;
      box-shadow: 0 12px 48px rgb(26 39 52 / 11%);

      .account-created {
        text-align: center;
        margin: 15px 0 40px;

        h1 {
          margin: 4px 0 0;
          color: #000;
          font-size: 20px;
          font-weight: 600;
        }

        p {
          font-size: 18px;
          margin: 5px 38px 0;
          line-height: 1.2;
        }

        svg {
          fill: #3391e1;
          height: 40px;
          stroke: #fff;
          stroke-width: 8px;
        }
      }

      .title {
        margin-top: 9px;
        margin-bottom: 27px;
        font-size: 22px;
        color: #111;
        word-spacing: -1px;
        font-weight: 500;
      }

      .i-have-an-account {
        position: absolute;
        top: 38px;
        right: 34px;
        font-size: 14px;
        border-bottom: 1.4px solid;
        padding-bottom: 2px;
        word-spacing: -1px;
        color: #0e71c3;
        text-decoration: none;

        &:hover {
          border-bottom: 2px solid;
        }
      }

      .form {
        position: relative;
        margin-bottom: 6px;
        padding-bottom: 11px;

        button#signup {
          width: 100%;
          margin-bottom: 5px;
          padding: 5px 0;
          height: 47px;
          font-size: 14px;
          border-radius: 4px !important;
        }

        .input-wrapper {
          position: relative;

          svg {
            position: absolute;
            right: 8.5px;
            top: 15.5px;
            width: 22px;
            height: 17px;
            cursor: pointer;
            padding: 3px;
          }

          svg path {
            fill: #646d75;
          }

          svg.revealed {
            top: 13.5px;
            height: 20px;
          }

          input {
            width: 385px;
            border: none;
            margin-bottom: 8px;
            padding: 14px 34px 0 12px;
            height: 46px;
            border-radius: 3px;
            font-size: 15px;
            transition:
              box-shadow 0.1s ease,
              border-color 0.25s ease;
            border: 1px solid #646b71;
          }

          input:focus ~ .placeholder,
          input:not(:focus).has-value ~ .placeholder {
            top: 7px;
            left: 12px;
            font-size: 11px;
          }

          .placeholder {
            position: absolute;
            pointer-events: none;
            left: 12px;
            top: 15px;
            color: #414954;
            font-size: 14px;
            transition: 0.2s ease all;
          }
        }

        .agree-to-terms {
          text-align: center;
          font-size: 12px;
          margin: 9px 0 11px;
          color: #586169;

          a {
            color: #586169;
            border-bottom: 1px solid #a1abb4;
            padding-bottom: 1px;
            text-decoration: none;

            &:hover {
              color: #2f3236;
              border-bottom: 1px solid #757d85;
            }
          }
        }
      }
    }

    .error {
      border: 1px solid #f6b4b4;
      padding: 11px 20px;
      background: #ffcccc70;
      color: #222;
      border-radius: 4px;
      font-weight: 400;
      font-size: 12px;
      position: relative;
      text-align: center;
      width: 100%;
      line-height: 1.3;
      margin: 5px auto 1px;

      button {
        background: none;
        border: none;
        padding: 0 0 3px;
        margin: 0;
        height: 13px;
        border-bottom: 1px solid #ea9898;
        color: #bb4343;
        font-size: 11.5px;
        cursor: pointer;

        &:hover {
          border-bottom: 1px solid #e47171;
          color: #bf2323;
        }
      }
    }
  </style>
</svelte:head>
