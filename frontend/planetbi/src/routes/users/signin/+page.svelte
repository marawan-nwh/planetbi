<script>
  import { title } from "$lib/js/stores.js";
  title.set("Sign in");

  import { hasValue } from "$lib/js/common.js";
  import { http } from "$lib/js/http.js";

  import { goto } from "$app/navigation";

  import Header from "$lib/components/auth-header/component.svelte";

  // state of the form
  // 0: initial
  // 1: email has been entered, ask for password
  // 2: password reset initialized, ask for token and new password
  // 3: password reset successfully
  let state = 0;

  let email;
  let password;
  let unverifiedEmail = false;

  let passwordResetToken;
  let passwordResetCorrelationToken = "";
  let newPassword;
  let newPasswordConfirmation;

  let UnverifiedEmailMessageID = 0;
  let SendingVerificationMessageID = 1;
  let VerificationEmailSentMessageID = 2;
  let VerificationEmailSentPreviouslyMessageID = 3;
  let messageID = UnverifiedEmailMessageID;

  let error = "";
  let disabled = false;
  let running = false;
  let revealPassword = false;

  function submitOnEnter(e) {
    if (e.keyCode === 13) {
      if (state === 0) {
        next();
      } else if (state === 1) {
        signin();
      } else if (state === 2) {
        resetPassword();
      }
    }
  }

  function next() {
    disabled = true;
    running = true;

    if (state === 0) {
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

      error = "";

      setTimeout(function () {
        state = 1;
        disabled = false;
        running = false;
      }, 500);
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

  function resendVerificationEmail() {
    error = "";
    disabled = true;

    // validate email
    if (!hasValue(email)) {
      error = "Please enter your email.";
      disabled = false;
      return;
    }
    if (!isValidEmail(email)) {
      error = "Please enter a valid email.";
      disabled = false;
      return;
    }

    // clear error
    error = "";

    messageID = SendingVerificationMessageID;

    http.post(
      "/resend-verification-email",
      { email },
      function (res, status) {
        messageID = VerificationEmailSentMessageID;
        disabled = false;
      },
      function (res, status) {
        if (status == 201) {
          messageID = VerificationEmailSentPreviouslyMessageID;
          disabled = false;
        } else {
          error = "Unexpected error, please try again.";
          disabled = false;
          messageID = UnverifiedEmailMessageID;
        }
      }
    );
  }

  function signin() {
    error = "";
    disabled = true;
    running = true;

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

    if (!hasValue(password)) {
      error = "Please enter your password.";
      disabled = false;
      running = false;
      return;
    }

    // clear error
    error = "";

    // signin
    http.post(
      "/signin",
      { email, password },
      function (res, status) {
        // redirect to homepage
        goto("/");
      },
      function (res, status) {
        if (status == 401) {
          error = "Invalid email or password.";
        } else if (status == 406) {
          messageID = UnverifiedEmailMessageID;
          unverifiedEmail = true;
        } else {
          error = "Unexpected error, please try again.";
        }
        disabled = false;
        running = false;
      }
    );
  }

  let initializing = false;
  function initResetPassword() {
    if (initializing) return;
    initializing = true;
    disabled = true;
    running = true;
    error = "";

    http.post(
      "/init-reset-password",
      { email },
      function (res, status) {
        passwordResetCorrelationToken = res.password_reset_correlation_token;

        initializing = false;
        disabled = false;
        running = false;
        state = 2;
        unverifiedEmail = false;
      },
      function (res, status) {
        error = "Unexpected error, please try again.";
        disabled = false;
        running = false;
        initializing = false;
      }
    );
  }

  function resetPassword() {
    error = "";
    disabled = true;
    running = true;

    // validation
    if (!hasValue(passwordResetToken)) {
      error = "Please enter the secret token.";
      disabled = false;
      running = false;
      return;
    }

    // validate password
    // don't add complexity requirements
    // https://learn.microsoft.com/en-us/windows/security/threat-protection/security-policy-settings/password-must-meet-complexity-requirements
    if (!hasValue(newPassword)) {
      error = "Please enter your password.";
      disabled = false;
      running = false;
      return;
    }
    if (newPassword.length < 8) {
      error = "Password must be 8+ chars.";
      disabled = false;
      running = false;
      return;
    }

    // confirm password
    if (newPassword !== newPasswordConfirmation) {
      error = "Passwords do not match.";
      disabled = false;
      running = false;
      return;
    }

    // clear error
    error = "";

    http.post(
      "/reset-password",
      {
        email,
        password_reset_correlation_token: passwordResetCorrelationToken,
        password_reset_token: passwordResetToken,
        new_password: newPassword,
      },
      function (res, status) {
        disabled = false;
        running = false;
        error = "";
        state = 3;
      },
      function (res, status) {
        error = "Unexpected error, please try again.";
        disabled = false;
        running = false;
      }
    );
  }
</script>

<Header />

<div class="form-wrapper">
  <div class="password-changed {state === 3 ? '' : 'hide'}">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      enable-background="new 0 0 256 256"
      viewBox="0 0 256 256"
    >
      <path
        fill-rule="evenodd"
        d="M27.364,99.627c3.874-4.29,10.742-4.86,15.268-1.26   l52.917,42.091c4.526,3.601,11.638,3.27,15.81-0.732L213.38,41.82c4.171-4.002,10.887-3.891,14.92,0.254l24.74,25.404   c4.036,4.139,3.931,10.809-0.236,14.814L116.078,213.814c-4.166,4.008-11.336,4.417-15.932,0.911L3.873,141.291   c-4.595-3.507-5.188-9.883-1.313-14.179L27.364,99.627z"
        clip-rule="evenodd"
      />
    </svg>
    <h1>Password changed!</h1>
    <p>
      Your password has been changed successfully.<br /><a
        href="/users/signin"
        on:click={() => {
          window.location.reload();
        }}>Sign in</a
      > with your new password.
    </p>
  </div>

  <div class="title {state === 3 ? 'hide' : ''}">
    {state === 0 || state === 1 ? "Sign in" : "Reset password"}
  </div>
  <a
    class="i-dont-have-an-account {state === 0 || state === 2 ? '' : 'hide'}"
    href="/users/signup">I don't have an account</a
  >

  <div
    class="forgot-password {state === 1 ? '' : 'hide'}"
    on:click={initResetPassword}
  >
    <span>Forgot password?</span>
  </div>

  <div class="secret-token-has-been-sent {state === 2 ? '' : 'hide'}">
    You should receive a secret token if the email you entered is registered.
    Please enter the token alongside the new password.
  </div>

  <div class="form">
    <div class="input-wrapper {state === 0 ? '' : 'hide'}">
      <input
        type="email"
        class={hasValue(email) ? "has-value" : ""}
        on:keydown={submitOnEnter}
        bind:value={email}
        {disabled}
      />
      <span class="placeholder">Email</span>
    </div>

    <div class="input-wrapper {state === 1 ? '' : 'hide'}">
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
          class="revealed {disabled ? 'disabled' : ''}"
          on:click={() => {
            if (!disabled) {
              revealPassword = false;
            }
          }}
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
          class={disabled ? "disabled" : ""}
          on:click={() => {
            if (!disabled) {
              revealPassword = true;
            }
          }}
          ><path
            d="M14.9421 5.61977L14.4918 5.4024V5.4024L14.9421 5.61977ZM1.05765 5.61977L1.50793 5.4024L1.50793 5.4024L1.05765 5.61977ZM14.9421 5.38039L15.3924 5.16301L15.3924 5.16301L14.9421 5.38039ZM1.05765 5.38039L0.607377 5.16301L0.607377 5.16301L1.05765 5.38039ZM14.4918 5.4024C11.8689 10.8355 4.1308 10.8355 1.50793 5.4024L0.607378 5.83715C3.59408 12.0239 12.4057 12.0239 15.3924 5.83715L14.4918 5.4024ZM14.4918 5.59776C14.462 5.53605 14.462 5.46411 14.4918 5.4024L15.3924 5.83715C15.4952 5.62419 15.4952 5.37596 15.3924 5.16301L14.4918 5.59776ZM1.50793 5.59776C4.1308 0.164679 11.8689 0.164678 14.4918 5.59776L15.3924 5.16301C12.4057 -1.02374 3.59408 -1.02374 0.607377 5.16301L1.50793 5.59776ZM0.607377 5.16301C0.504572 5.37597 0.504572 5.6242 0.607378 5.83715L1.50793 5.4024C1.53772 5.46411 1.53772 5.53605 1.50793 5.59776L0.607377 5.16301ZM9.99987 5.50008C9.99987 6.60465 9.10444 7.50008 7.99987 7.50008V8.50008C9.65672 8.50008 10.9999 7.15693 10.9999 5.50008H9.99987ZM7.99987 7.50008C6.8953 7.50008 5.99987 6.60465 5.99987 5.50008H4.99987C4.99987 7.15693 6.34302 8.50008 7.99987 8.50008V7.50008ZM5.99987 5.50008C5.99987 4.39551 6.8953 3.50008 7.99987 3.50008V2.50008C6.34302 2.50008 4.99987 3.84322 4.99987 5.50008H5.99987ZM7.99987 3.50008C9.10444 3.50008 9.99987 4.39551 9.99987 5.50008H10.9999C10.9999 3.84322 9.65672 2.50008 7.99987 2.50008V3.50008Z"
          ></path></svg
        >
      {/if}
    </div>

    <div class="input-wrapper {state === 2 ? '' : 'hide'}">
      <input
        type="text"
        class={hasValue(passwordResetToken) ? "has-value" : ""}
        on:keydown={submitOnEnter}
        bind:value={passwordResetToken}
        {disabled}
      />
      <span class="placeholder">Secret token</span>
    </div>

    <div class="input-wrapper {state === 2 ? '' : 'hide'}">
      {#if revealPassword}
        <input
          type="text"
          class={hasValue(newPassword) ? "has-value" : ""}
          on:keydown={submitOnEnter}
          bind:value={newPassword}
          {disabled}
        />
        <span class="placeholder">New password</span>

        <svg
          viewBox="0 0 16 14"
          fill="none"
          aria-hidden="true"
          class="revealed {disabled ? 'disabled' : ''}"
          on:click={() => {
            if (!disabled) {
              revealPassword = false;
            }
          }}
          ><path
            d="M14.9421 7.6197L14.4918 7.40232V7.40232L14.9421 7.6197ZM1.05765 7.6197L1.50793 7.40232L1.50793 7.40232L1.05765 7.6197ZM14.9421 7.38031L15.3924 7.16294L15.3924 7.16293L14.9421 7.38031ZM1.05765 7.38031L0.607377 7.16293L0.607377 7.16294L1.05765 7.38031ZM10.9999 7.5C10.9999 7.22386 10.776 7 10.4999 7C10.2237 7 9.99987 7.22386 9.99987 7.5H10.9999ZM7.99987 9.5C7.72373 9.5 7.49987 9.72386 7.49987 10C7.49987 10.2761 7.72373 10.5 7.99987 10.5V9.5ZM1.14632 13.1464C0.951062 13.3417 0.951062 13.6583 1.14632 13.8536C1.34159 14.0488 1.65817 14.0488 1.85343 13.8536L1.14632 13.1464ZM14.8534 0.853553C15.0487 0.658291 15.0487 0.341709 14.8534 0.146447C14.6582 -0.0488155 14.3416 -0.0488155 14.1463 0.146447L14.8534 0.853553ZM11.0879 4.20226C11.3382 4.31901 11.6357 4.21079 11.7524 3.96055C11.8692 3.7103 11.761 3.41279 11.5107 3.29603L11.0879 4.20226ZM13.3114 4.45031C13.1009 4.2715 12.7854 4.29713 12.6066 4.50756C12.4278 4.71799 12.4534 5.03353 12.6638 5.21234L13.3114 4.45031ZM3.78294 11.331C4.0204 11.4719 4.32716 11.3937 4.46811 11.1562C4.60906 10.9187 4.53082 10.612 4.29335 10.471L3.78294 11.331ZM6.12685 11.2361C5.85976 11.166 5.58639 11.3257 5.51627 11.5928C5.44614 11.8599 5.60582 12.1332 5.87291 12.2033L6.12685 11.2361ZM5.5293 9.20225C5.68614 9.42953 5.99753 9.48663 6.22481 9.32978C6.45209 9.17294 6.50918 8.86155 6.35234 8.63427L5.5293 9.20225ZM9.09968 5.82922C9.33016 5.98132 9.6403 5.91778 9.7924 5.6873C9.9445 5.45682 9.88096 5.14668 9.65048 4.99458L9.09968 5.82922ZM14.4918 7.59768C14.462 7.53597 14.462 7.46404 14.4918 7.40232L15.3924 7.83707C15.4952 7.62411 15.4952 7.37589 15.3924 7.16294L14.4918 7.59768ZM0.607377 7.16294C0.504572 7.37589 0.504572 7.62412 0.607378 7.83707L1.50793 7.40232C1.53772 7.46403 1.53772 7.53597 1.50793 7.59768L0.607377 7.16294ZM9.99987 7.5C9.99987 8.60457 9.10444 9.5 7.99987 9.5V10.5C9.65672 10.5 10.9999 9.15686 10.9999 7.5H9.99987ZM5.99987 7.5C5.99987 6.39543 6.8953 5.5 7.99987 5.5V4.5C6.34302 4.5 4.99987 5.84315 4.99987 7.5H5.99987ZM1.85343 13.8536L14.8534 0.853553L14.1463 0.146447L1.14632 13.1464L1.85343 13.8536ZM1.50793 7.59768C3.35244 3.77692 7.73141 2.63626 11.0879 4.20226L11.5107 3.29603C7.70772 1.52172 2.71273 2.80184 0.607377 7.16293L1.50793 7.59768ZM12.6638 5.21234C13.3942 5.83295 14.0231 6.62685 14.4918 7.59768L15.3924 7.16293C14.8632 6.06675 14.1481 5.16133 13.3114 4.45031L12.6638 5.21234ZM4.29335 10.471C3.15305 9.79421 2.17002 8.7738 1.50793 7.40232L0.607378 7.83707C1.35622 9.38825 2.47624 10.5554 3.78294 11.331L4.29335 10.471ZM14.4918 7.40232C12.8631 10.7762 9.26241 12.0594 6.12685 11.2361L5.87291 12.2033C9.42276 13.1354 13.5347 11.6852 15.3924 7.83707L14.4918 7.40232ZM6.35234 8.63427C6.13001 8.31211 5.99987 7.92201 5.99987 7.5H4.99987C4.99987 8.13132 5.19543 8.71846 5.5293 9.20225L6.35234 8.63427ZM7.99987 5.5C8.40703 5.5 8.78437 5.62114 9.09968 5.82922L9.65048 4.99458C9.17675 4.68196 8.6088 4.5 7.99987 4.5V5.5Z"
            fill="#6D758D"
          ></path></svg
        >
      {:else}
        <input
          type="password"
          class={hasValue(newPassword) ? "has-value" : ""}
          on:keydown={submitOnEnter}
          bind:value={newPassword}
          {disabled}
        />
        <span class="placeholder">New password</span>

        <svg
          viewBox="0 0 16 11"
          fill="none"
          aria-hidden="true"
          class={disabled ? "disabled" : ""}
          on:click={() => {
            if (!disabled) {
              revealPassword = true;
            }
          }}
          ><path
            d="M14.9421 5.61977L14.4918 5.4024V5.4024L14.9421 5.61977ZM1.05765 5.61977L1.50793 5.4024L1.50793 5.4024L1.05765 5.61977ZM14.9421 5.38039L15.3924 5.16301L15.3924 5.16301L14.9421 5.38039ZM1.05765 5.38039L0.607377 5.16301L0.607377 5.16301L1.05765 5.38039ZM14.4918 5.4024C11.8689 10.8355 4.1308 10.8355 1.50793 5.4024L0.607378 5.83715C3.59408 12.0239 12.4057 12.0239 15.3924 5.83715L14.4918 5.4024ZM14.4918 5.59776C14.462 5.53605 14.462 5.46411 14.4918 5.4024L15.3924 5.83715C15.4952 5.62419 15.4952 5.37596 15.3924 5.16301L14.4918 5.59776ZM1.50793 5.59776C4.1308 0.164679 11.8689 0.164678 14.4918 5.59776L15.3924 5.16301C12.4057 -1.02374 3.59408 -1.02374 0.607377 5.16301L1.50793 5.59776ZM0.607377 5.16301C0.504572 5.37597 0.504572 5.6242 0.607378 5.83715L1.50793 5.4024C1.53772 5.46411 1.53772 5.53605 1.50793 5.59776L0.607377 5.16301ZM9.99987 5.50008C9.99987 6.60465 9.10444 7.50008 7.99987 7.50008V8.50008C9.65672 8.50008 10.9999 7.15693 10.9999 5.50008H9.99987ZM7.99987 7.50008C6.8953 7.50008 5.99987 6.60465 5.99987 5.50008H4.99987C4.99987 7.15693 6.34302 8.50008 7.99987 8.50008V7.50008ZM5.99987 5.50008C5.99987 4.39551 6.8953 3.50008 7.99987 3.50008V2.50008C6.34302 2.50008 4.99987 3.84322 4.99987 5.50008H5.99987ZM7.99987 3.50008C9.10444 3.50008 9.99987 4.39551 9.99987 5.50008H10.9999C10.9999 3.84322 9.65672 2.50008 7.99987 2.50008V3.50008Z"
          ></path></svg
        >
      {/if}
    </div>

    <div class="input-wrapper {state === 2 ? '' : 'hide'}">
      {#if revealPassword}
        <input
          type="text"
          class={hasValue(newPasswordConfirmation) ? "has-value" : ""}
          on:keydown={submitOnEnter}
          bind:value={newPasswordConfirmation}
          {disabled}
        />
        <span class="placeholder">Confirm password</span>

        <svg
          viewBox="0 0 16 14"
          fill="none"
          aria-hidden="true"
          class="revealed {disabled ? 'disabled' : ''}"
          on:click={() => {
            if (!disabled) {
              revealPassword = false;
            }
          }}
          ><path
            d="M14.9421 7.6197L14.4918 7.40232V7.40232L14.9421 7.6197ZM1.05765 7.6197L1.50793 7.40232L1.50793 7.40232L1.05765 7.6197ZM14.9421 7.38031L15.3924 7.16294L15.3924 7.16293L14.9421 7.38031ZM1.05765 7.38031L0.607377 7.16293L0.607377 7.16294L1.05765 7.38031ZM10.9999 7.5C10.9999 7.22386 10.776 7 10.4999 7C10.2237 7 9.99987 7.22386 9.99987 7.5H10.9999ZM7.99987 9.5C7.72373 9.5 7.49987 9.72386 7.49987 10C7.49987 10.2761 7.72373 10.5 7.99987 10.5V9.5ZM1.14632 13.1464C0.951062 13.3417 0.951062 13.6583 1.14632 13.8536C1.34159 14.0488 1.65817 14.0488 1.85343 13.8536L1.14632 13.1464ZM14.8534 0.853553C15.0487 0.658291 15.0487 0.341709 14.8534 0.146447C14.6582 -0.0488155 14.3416 -0.0488155 14.1463 0.146447L14.8534 0.853553ZM11.0879 4.20226C11.3382 4.31901 11.6357 4.21079 11.7524 3.96055C11.8692 3.7103 11.761 3.41279 11.5107 3.29603L11.0879 4.20226ZM13.3114 4.45031C13.1009 4.2715 12.7854 4.29713 12.6066 4.50756C12.4278 4.71799 12.4534 5.03353 12.6638 5.21234L13.3114 4.45031ZM3.78294 11.331C4.0204 11.4719 4.32716 11.3937 4.46811 11.1562C4.60906 10.9187 4.53082 10.612 4.29335 10.471L3.78294 11.331ZM6.12685 11.2361C5.85976 11.166 5.58639 11.3257 5.51627 11.5928C5.44614 11.8599 5.60582 12.1332 5.87291 12.2033L6.12685 11.2361ZM5.5293 9.20225C5.68614 9.42953 5.99753 9.48663 6.22481 9.32978C6.45209 9.17294 6.50918 8.86155 6.35234 8.63427L5.5293 9.20225ZM9.09968 5.82922C9.33016 5.98132 9.6403 5.91778 9.7924 5.6873C9.9445 5.45682 9.88096 5.14668 9.65048 4.99458L9.09968 5.82922ZM14.4918 7.59768C14.462 7.53597 14.462 7.46404 14.4918 7.40232L15.3924 7.83707C15.4952 7.62411 15.4952 7.37589 15.3924 7.16294L14.4918 7.59768ZM0.607377 7.16294C0.504572 7.37589 0.504572 7.62412 0.607378 7.83707L1.50793 7.40232C1.53772 7.46403 1.53772 7.53597 1.50793 7.59768L0.607377 7.16294ZM9.99987 7.5C9.99987 8.60457 9.10444 9.5 7.99987 9.5V10.5C9.65672 10.5 10.9999 9.15686 10.9999 7.5H9.99987ZM5.99987 7.5C5.99987 6.39543 6.8953 5.5 7.99987 5.5V4.5C6.34302 4.5 4.99987 5.84315 4.99987 7.5H5.99987ZM1.85343 13.8536L14.8534 0.853553L14.1463 0.146447L1.14632 13.1464L1.85343 13.8536ZM1.50793 7.59768C3.35244 3.77692 7.73141 2.63626 11.0879 4.20226L11.5107 3.29603C7.70772 1.52172 2.71273 2.80184 0.607377 7.16293L1.50793 7.59768ZM12.6638 5.21234C13.3942 5.83295 14.0231 6.62685 14.4918 7.59768L15.3924 7.16293C14.8632 6.06675 14.1481 5.16133 13.3114 4.45031L12.6638 5.21234ZM4.29335 10.471C3.15305 9.79421 2.17002 8.7738 1.50793 7.40232L0.607378 7.83707C1.35622 9.38825 2.47624 10.5554 3.78294 11.331L4.29335 10.471ZM14.4918 7.40232C12.8631 10.7762 9.26241 12.0594 6.12685 11.2361L5.87291 12.2033C9.42276 13.1354 13.5347 11.6852 15.3924 7.83707L14.4918 7.40232ZM6.35234 8.63427C6.13001 8.31211 5.99987 7.92201 5.99987 7.5H4.99987C4.99987 8.13132 5.19543 8.71846 5.5293 9.20225L6.35234 8.63427ZM7.99987 5.5C8.40703 5.5 8.78437 5.62114 9.09968 5.82922L9.65048 4.99458C9.17675 4.68196 8.6088 4.5 7.99987 4.5V5.5Z"
            fill="#6D758D"
          ></path></svg
        >
      {:else}
        <input
          type="password"
          class={hasValue(newPasswordConfirmation) ? "has-value" : ""}
          on:keydown={submitOnEnter}
          bind:value={newPasswordConfirmation}
          {disabled}
        />
        <span class="placeholder">Confirm password</span>

        <svg
          viewBox="0 0 16 11"
          fill="none"
          aria-hidden="true"
          class={disabled ? "disabled" : ""}
          on:click={() => {
            if (!disabled) {
              revealPassword = true;
            }
          }}
          ><path
            d="M14.9421 5.61977L14.4918 5.4024V5.4024L14.9421 5.61977ZM1.05765 5.61977L1.50793 5.4024L1.50793 5.4024L1.05765 5.61977ZM14.9421 5.38039L15.3924 5.16301L15.3924 5.16301L14.9421 5.38039ZM1.05765 5.38039L0.607377 5.16301L0.607377 5.16301L1.05765 5.38039ZM14.4918 5.4024C11.8689 10.8355 4.1308 10.8355 1.50793 5.4024L0.607378 5.83715C3.59408 12.0239 12.4057 12.0239 15.3924 5.83715L14.4918 5.4024ZM14.4918 5.59776C14.462 5.53605 14.462 5.46411 14.4918 5.4024L15.3924 5.83715C15.4952 5.62419 15.4952 5.37596 15.3924 5.16301L14.4918 5.59776ZM1.50793 5.59776C4.1308 0.164679 11.8689 0.164678 14.4918 5.59776L15.3924 5.16301C12.4057 -1.02374 3.59408 -1.02374 0.607377 5.16301L1.50793 5.59776ZM0.607377 5.16301C0.504572 5.37597 0.504572 5.6242 0.607378 5.83715L1.50793 5.4024C1.53772 5.46411 1.53772 5.53605 1.50793 5.59776L0.607377 5.16301ZM9.99987 5.50008C9.99987 6.60465 9.10444 7.50008 7.99987 7.50008V8.50008C9.65672 8.50008 10.9999 7.15693 10.9999 5.50008H9.99987ZM7.99987 7.50008C6.8953 7.50008 5.99987 6.60465 5.99987 5.50008H4.99987C4.99987 7.15693 6.34302 8.50008 7.99987 8.50008V7.50008ZM5.99987 5.50008C5.99987 4.39551 6.8953 3.50008 7.99987 3.50008V2.50008C6.34302 2.50008 4.99987 3.84322 4.99987 5.50008H5.99987ZM7.99987 3.50008C9.10444 3.50008 9.99987 4.39551 9.99987 5.50008H10.9999C10.9999 3.84322 9.65672 2.50008 7.99987 2.50008V3.50008Z"
          ></path></svg
        >
      {/if}
    </div>

    <button
      class="next btn-blue ld-over {running ? 'running' : ''} {state === 0
        ? ''
        : 'hide'}"
      id="next"
      on:click={next}
      {disabled}
    >
      <div class="ld ld-ring ld-cycle"></div>
      Next
    </button>

    <button
      class="signin btn-blue ld-over {running ? 'running' : ''} {state === 1
        ? ''
        : 'hide'}"
      id="signin"
      on:click={signin}
      {disabled}
    >
      <div class="ld ld-ring ld-cycle"></div>
      Sign in
    </button>

    <button
      class="reset-password btn-blue ld-over {running
        ? 'running'
        : ''} {state === 2 ? '' : 'hide'}"
      id="reset-password"
      on:click={resetPassword}
      {disabled}
    >
      <div class="ld ld-ring ld-cycle"></div>
      Reset password
    </button>

    <div class="error {!error ? 'hide' : ''}">
      <span>{error}</span>
      <button
        on:click={() => {
          error = "";
        }}>hide</button
      >
    </div>

    <div class="message {unverifiedEmail ? '' : 'hide'}">
      {#if messageID == UnverifiedEmailMessageID}
        <span>Unverified email.</span>
        <button on:click={resendVerificationEmail}
          >Resend verification email</button
        >
      {:else if messageID == SendingVerificationMessageID}
        <span>Sending...</span>
      {:else if messageID == VerificationEmailSentMessageID}
        <span>New verification email has been sent!</span>
        <button
          on:click={() => {
            unverifiedEmail = false;
            messageID = UnverifiedEmailMessageID;
          }}>hide</button
        >
      {:else if messageID == VerificationEmailSentPreviouslyMessageID}
        <span>Already sent, try again after 5 minutes.</span>
        <button
          on:click={() => {
            unverifiedEmail = false;
            messageID = UnverifiedEmailMessageID;
          }}>hide</button
        >
      {/if}
    </div>
  </div>
</div>

<svelte:head>
  <style>
    body {
      background: #f2f7fc;
    }
  </style>
</svelte:head>

<style src="./page.css" lang="scss"></style>
