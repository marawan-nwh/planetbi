<script>
  import { title } from "$lib/stores.js";
  title.set("Email Verification");
</script>

<div class="verification-wrapper">
  <div class="verified-already hide">
    <div class="already-verified-title">Email already verified</div>
    <div class="already-verified-button">
      <a href="/signin" class="btn-blue">Signin</a>
    </div>
  </div>

  <div class="verification-error hide">
    <div class="verification-error-title">Verification error</div>
    <div class="verification-error-text">
      Something went wrong, please try again.
    </div>
  </div>

  <div class="verification-sucess hide">
    <div class="verification-sucess-title">Verified successfully</div>
    <div class="verification-sucess-button">
      <a href="/signin" class="btn-blue">Signin</a>
    </div>
  </div>

  <div class="verification">
    <div class="userid">User ID</div>
    <div id="userid"></div>
    <button class="btn-blue ld-over" onclick="verifyEmail(event)">
      <div class="ld ld-ring ld-cycle"></div>
      Verify my email
    </button>
  </div>
</div>

<style>
  @import "verify.css";
</style>
