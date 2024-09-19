<script>
  import { title } from "$lib/js/stores.js";
  title.set("Email Verification");

  import { hasValue } from "$lib/js/common.js";
  import { http } from "$lib/js/http.js";
  import { onMount } from "svelte";

  import Header from "$lib/components/auth-header/component.svelte";

  let userID;
  let token;

  onMount(() => {
    const url = new URL(window.location.href);
    userID = url.searchParams.get("user_id");
    token = url.searchParams.get("token");
    if (!hasValue(userID) || !hasValue(token)) {
      error = "Invalid verification link.";
      disabled = true;
    }
  });

  let success = false;
  let disabled = false;
  let running = false;
  let error = "";

  function verify() {
    disabled = true;
    running = true;

    http.post(
      "/verify",
      { user_id: userID, token: token },
      function (res, status) {
        error = "";
        success = true;
        disabled = false;
        running = false;
      },
      function (res, status) {
        if (status === 401) {
          error = "Invalid verification link.";
        } else {
          error = "Something went wrong, please try again.";
        }
        disabled = false;
        running = false;
      }
    );
  }
</script>

<Header />

<div class="verification-wrapper">
  <div class="sucess {success ? '' : 'hide'}">
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
    <h1>Verified successfully!</h1>
    <p>
      You can now <a href="/users/signin">sign in</a>, and start using PlanetBI.
    </p>
  </div>

  <div class="error {hasValue(error) ? '' : 'hide'}">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      xmlns:xlink="http://www.w3.org/1999/xlink"
      height="40"
      width="40"
      viewBox="0 0 512 512"
      xml:space="preserve"
    >
      <g>
        <g>
          <g>
            <path
              d="M437.016,74.984c-99.979-99.979-262.075-99.979-362.033,0.002c-99.978,99.978-99.978,262.073,0.004,362.031     c99.954,99.978,262.05,99.978,362.029-0.002C536.995,337.059,536.995,174.964,437.016,74.984z M406.848,406.844     c-83.318,83.318-218.396,83.318-301.691,0.004c-83.318-83.299-83.318-218.377-0.002-301.693     c83.297-83.317,218.375-83.317,301.691,0S490.162,323.549,406.848,406.844z"
            />
            <path
              d="M361.592,150.408c-8.331-8.331-21.839-8.331-30.17,0l-75.425,75.425l-75.425-75.425c-8.331-8.331-21.839-8.331-30.17,0     s-8.331,21.839,0,30.17l75.425,75.425L150.43,331.4c-8.331,8.331-8.331,21.839,0,30.17c8.331,8.331,21.839,8.331,30.17,0     l75.397-75.397l75.419,75.419c8.331,8.331,21.839,8.331,30.17,0c8.331-8.331,8.331-21.839,0-30.17l-75.419-75.419l75.425-75.425     C369.923,172.247,369.923,158.74,361.592,150.408z"
            />
          </g>
        </g>
      </g>
    </svg>

    <h1>Error</h1>
    <p>
      {error}
    </p>
  </div>

  <div class="verification {success || hasValue(error) ? 'hide' : ''}">
    <div class="title">Email verification</div>
    <div class="agree-to-terms">
      By signing up, you agree to the <a href="/terms-and-conditions"
        >Terms and Conditions</a
      >.
    </div>
    <button
      class="btn-blue ld-over {running ? 'running' : ''}"
      on:click={verify}
      {disabled}
    >
      <div class="ld ld-ring ld-cycle"></div>
      Verify my email
    </button>
  </div>
</div>

<svelte:head>
  <style>
    body {
      position: relative;
      overflow: hidden;
      margin-top: 0;
      padding-top: 0;
      background: #f2f7fc;
    }

    .verification-wrapper {
      margin: 44px auto 45px;
      width: 450px;
      position: relative;
      background: #fff;
      border-radius: 8px;
      padding: 25px 32px 26px;
      box-shadow: 0 12px 48px rgb(26 39 52 / 11%);

      .verification {
        .title {
          margin-top: 9px;
          margin-bottom: 21px;
          font-size: 22px;
          color: #111;
          word-spacing: -1px;
          font-weight: 500;
        }

        button {
          width: 100%;
          margin-bottom: 5px;
          padding: 5px 0;
          height: 47px;
          font-size: 14px;
        }
      }
    }

    .sucess,
    .error {
      text-align: center;
      margin: 15px 0 40px;

      h1 {
        margin: 4px 0 0;
        color: #000;
        font-size: 20px;
        font-weight: 600;
      }

      p {
        font-size: 17px;
        margin: 5px 0px 0;
        line-height: 1.2;
      }

      a {
        color: #0e71c3;
        text-decoration: none;
      }

      a:hover {
        color: #0f60a3;
        text-decoration: underline;
      }
    }

    .sucess {
      svg {
        fill: #3391e1;
        width: 40px;
      }
    }

    .error {
      svg {
        fill: #e35757;
      }
    }

    .agree-to-terms {
      text-align: center;
      font-size: 12px;
      margin: 9px 0 11px;
      color: #52606c;

      a {
        color: #646d75;
        border-bottom: 1px solid #a1abb4;
        padding-bottom: 1px;
        text-decoration: none;

        &:hover {
          color: #2f3236;
          border-bottom: 1px solid #757d85;
        }
      }
    }
  </style>
</svelte:head>
