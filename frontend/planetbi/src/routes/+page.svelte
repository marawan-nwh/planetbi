<script>
  import { title } from "$lib/js/stores.js";
  title.set("");

  import Header from "$lib/components/header/component.svelte";
</script>

<Header />

<sidebar>
  <div class="nav">
    <div class="title">Home</div>
    <div class="item active item-all">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M3.75 6A2.25 2.25 0 0 1 6 3.75h2.25A2.25 2.25 0 0 1 10.5 6v2.25a2.25 2.25 0 0 1-2.25 2.25H6a2.25 2.25 0 0 1-2.25-2.25V6ZM3.75 15.75A2.25 2.25 0 0 1 6 13.5h2.25a2.25 2.25 0 0 1 2.25 2.25V18a2.25 2.25 0 0 1-2.25 2.25H6A2.25 2.25 0 0 1 3.75 18v-2.25ZM13.5 6a2.25 2.25 0 0 1 2.25-2.25H18A2.25 2.25 0 0 1 20.25 6v2.25A2.25 2.25 0 0 1 18 10.5h-2.25a2.25 2.25 0 0 1-2.25-2.25V6ZM13.5 15.75a2.25 2.25 0 0 1 2.25-2.25H18a2.25 2.25 0 0 1 2.25 2.25V18A2.25 2.25 0 0 1 18 20.25h-2.25A2.25 2.25 0 0 1 13.5 18v-2.25Z"
        ></path>
      </svg>

      All
    </div>
    <div class="item item-starred">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M11.48 3.499a.562.562 0 0 1 1.04 0l2.125 5.111a.563.563 0 0 0 .475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 0 0-.182.557l1.285 5.385a.562.562 0 0 1-.84.61l-4.725-2.885a.562.562 0 0 0-.586 0L6.982 20.54a.562.562 0 0 1-.84-.61l1.285-5.386a.562.562 0 0 0-.182-.557l-4.204-3.602a.562.562 0 0 1 .321-.988l5.518-.442a.563.563 0 0 0 .475-.345L11.48 3.5Z"
        ></path>
      </svg>

      Starred
    </div>
    <div class="item item-recent">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M12 6v6h4.5m4.5 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
        ></path>
      </svg>

      Recent
    </div>
    <div class="item item-archive">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="m20.25 7.5-.625 10.632a2.25 2.25 0 0 1-2.247 2.118H6.622a2.25 2.25 0 0 1-2.247-2.118L3.75 7.5M10 11.25h4M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125Z"
        ></path>
      </svg>

      Archive
    </div>
  </div>

  <div class="folders">
    <div class="title">Folders</div>
    <div class="folder">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M2.25 12.75V12A2.25 2.25 0 0 1 4.5 9.75h15A2.25 2.25 0 0 1 21.75 12v.75m-8.69-6.44-2.12-2.12a1.5 1.5 0 0 0-1.061-.44H4.5A2.25 2.25 0 0 0 2.25 6v12a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9a2.25 2.25 0 0 0-2.25-2.25h-5.379a1.5 1.5 0 0 1-1.06-.44Z"
        ></path>
      </svg>

      Finance
    </div>
    <div class="folder">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M2.25 12.75V12A2.25 2.25 0 0 1 4.5 9.75h15A2.25 2.25 0 0 1 21.75 12v.75m-8.69-6.44-2.12-2.12a1.5 1.5 0 0 0-1.061-.44H4.5A2.25 2.25 0 0 0 2.25 6v12a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9a2.25 2.25 0 0 0-2.25-2.25h-5.379a1.5 1.5 0 0 1-1.06-.44Z"
        ></path>
      </svg>

      Engineering
    </div>
    <div class="new"></div>
  </div>

  <div class="trial-ends-soon">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="1.5"
      stroke="currentColor"
      class="w-6 h-6"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 0 1 2.25-2.25h13.5A2.25 2.25 0 0 1 21 7.5v11.25m-18 0A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75m-18 0v-7.5A2.25 2.25 0 0 1 5.25 9h13.5A2.25 2.25 0 0 1 21 11.25v7.5"
      ></path>
    </svg>

    Trial ends in 10 days

    <div class="upgrade">Upgrade</div>
  </div>
</sidebar>

<dashboards>
  <div class="header">
    <div class="title">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M3.75 6A2.25 2.25 0 0 1 6 3.75h2.25A2.25 2.25 0 0 1 10.5 6v2.25a2.25 2.25 0 0 1-2.25 2.25H6a2.25 2.25 0 0 1-2.25-2.25V6ZM3.75 15.75A2.25 2.25 0 0 1 6 13.5h2.25a2.25 2.25 0 0 1 2.25 2.25V18a2.25 2.25 0 0 1-2.25 2.25H6A2.25 2.25 0 0 1 3.75 18v-2.25ZM13.5 6a2.25 2.25 0 0 1 2.25-2.25H18A2.25 2.25 0 0 1 20.25 6v2.25A2.25 2.25 0 0 1 18 10.5h-2.25a2.25 2.25 0 0 1-2.25-2.25V6ZM13.5 15.75a2.25 2.25 0 0 1 2.25-2.25H18a2.25 2.25 0 0 1 2.25 2.25V18A2.25 2.25 0 0 1 18 20.25h-2.25A2.25 2.25 0 0 1 13.5 18v-2.25Z"
        ></path>
      </svg>

      All
    </div>
    <div class="actions">
      <div class="filter">
        <svg role="graphics-symbol" viewBox="0 0 16 16">
          <path
            fill-rule="evenodd"
            clip-rule="evenodd"
            d="M2 4C2 3.58579 2.33579 3.25 2.75 3.25H13.25C13.6642 3.25 14 3.58579 14 4C14 4.41421 13.6642 4.75 13.25 4.75H2.75C2.33579 4.75 2 4.41421 2 4Z"
          ></path>
          <path
            fill-rule="evenodd"
            clip-rule="evenodd"
            d="M3.75 8C3.75 7.58579 4.08579 7.25 4.5 7.25H11.5C11.9142 7.25 12.25 7.58579 12.25 8C12.25 8.41421 11.9142 8.75 11.5 8.75H4.5C4.08579 8.75 3.75 8.41421 3.75 8Z"
          ></path>
          <path
            fill-rule="evenodd"
            clip-rule="evenodd"
            d="M5.5 12C5.5 11.5858 5.83579 11.25 6.25 11.25H9.75C10.1642 11.25 10.5 11.5858 10.5 12C10.5 12.4142 10.1642 12.75 9.75 12.75H6.25C5.83579 12.75 5.5 12.4142 5.5 12Z"
          ></path>
        </svg>
        Filter...
      </div>
      <div class="create">
        Create
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 16 16"
          fill="currentColor"
          class="w-4 h-4"
        >
          <path
            fill-rule="evenodd"
            d="M4.22 6.22a.75.75 0 0 1 1.06 0L8 8.94l2.72-2.72a.75.75 0 1 1 1.06 1.06l-3.25 3.25a.75.75 0 0 1-1.06 0L4.22 7.28a.75.75 0 0 1 0-1.06Z"
            clip-rule="evenodd"
          ></path>
        </svg>
      </div>
    </div>
  </div>
  <div class="cls">
    <div class="cl title">Name</div>
    <div class="cl type">Type</div>
    <div class="cl visibility">Visibility</div>
    <div class="cl created">Created</div>
  </div>
  <div class="list">
    <div class="item">
      <a href="/dashboards?id=ckcp185ucdm4n11k36s0">
        <div class="title">Sales 🤑</div>
        <div class="type">Dashboard</div>
        <div class="visibility">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="13"
            height="13"
            viewBox="0 0 24 24"
            fill="none"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="tabler-icon tabler-icon-users"
          >
            <path d="M9 7m-4 0a4 4 0 1 0 8 0a4 4 0 1 0 -8 0"></path>
            <path d="M3 21v-2a4 4 0 0 1 4 -4h4a4 4 0 0 1 4 4v2"></path>
            <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
            <path d="M21 21v-2a4 4 0 0 0 -3 -3.85"></path>
          </svg>
          Shared
        </div>
        <div class="created">7 months ago</div>
      </a>
      <div onclick="dropdown(this)" class="actions dropbtn">
        <svg
          class="dropbtn"
          width="16"
          height="16"
          viewBox="0 0 16 16"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            class="dropbtn"
            d="M3 6.5C2.17157 6.5 1.5 7.17157 1.5 8C1.5 8.82843 2.17157 9.5 3 9.5C3.82843 9.5 4.5 8.82843 4.5 8C4.5 7.17157 3.82843 6.5 3 6.5Z"
          ></path>
          <path
            class="dropbtn"
            d="M6.5 8C6.5 7.17157 7.17157 6.5 8 6.5C8.82843 6.5 9.5 7.17157 9.5 8C9.5 8.82843 8.82843 9.5 8 9.5C7.17157 9.5 6.5 8.82843 6.5 8Z"
          ></path>
          <path
            class="dropbtn"
            d="M13 6.5C12.1716 6.5 11.5 7.17157 11.5 8C11.5 8.82843 12.1716 9.5 13 9.5C13.8284 9.5 14.5 8.82843 14.5 8C14.5 7.17157 13.8284 6.5 13 6.5Z"
          ></path>
        </svg>
        <div class="dropdown list-dropdown" style="right: 0">
          <div
            id="dashboard-list-dropdown-ckcp185ucdm4n11k36s0"
            class="dropdown-content dashboards-list-dropdown-content"
          >
            <a href="#" class="rename-dashboard-btn">Rename</a>
            <a href="#" class="delete-dashboard-btn">Delete</a>
          </div>
        </div>
      </div>
    </div>

    <div class="item _monitor">
      <a href="/viewsets?id=ckcp1htucdm4n11k36sg">
        <div class="title">Marketing</div>
        <div class="type">Viewset</div>
        <div class="visibility">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="13.5px"
            height="13.5px"
            viewBox="0 0 24 24"
            fill="none"
            style="left: -1px"
          >
            <path
              d="M7 10.0288C7.47142 10 8.05259 10 8.8 10H15.2C15.9474 10 16.5286 10 17 10.0288M7 10.0288C6.41168 10.0647 5.99429 10.1455 5.63803 10.327C5.07354 10.6146 4.6146 11.0735 4.32698 11.638C4 12.2798 4 13.1198 4 14.8V16.2C4 17.8802 4 18.7202 4.32698 19.362C4.6146 19.9265 5.07354 20.3854 5.63803 20.673C6.27976 21 7.11984 21 8.8 21H15.2C16.8802 21 17.7202 21 18.362 20.673C18.9265 20.3854 19.3854 19.9265 19.673 19.362C20 18.7202 20 17.8802 20 16.2V14.8C20 13.1198 20 12.2798 19.673 11.638C19.3854 11.0735 18.9265 10.6146 18.362 10.327C18.0057 10.1455 17.5883 10.0647 17 10.0288M7 10.0288V8C7 5.23858 9.23858 3 12 3C14.7614 3 17 5.23858 17 8V10.0288"
              stroke-width="1.8"
              stroke-linecap="round"
              stroke-linejoin="round"
            ></path>
          </svg>
          private
        </div>
        <div class="created">7 months ago</div>
      </a>
      <div onclick="dropdown(this)" class="actions dropbtn">
        <svg
          class="dropbtn"
          width="16"
          height="16"
          viewBox="0 0 16 16"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            class="dropbtn"
            d="M3 6.5C2.17157 6.5 1.5 7.17157 1.5 8C1.5 8.82843 2.17157 9.5 3 9.5C3.82843 9.5 4.5 8.82843 4.5 8C4.5 7.17157 3.82843 6.5 3 6.5Z"
          ></path>
          <path
            class="dropbtn"
            d="M6.5 8C6.5 7.17157 7.17157 6.5 8 6.5C8.82843 6.5 9.5 7.17157 9.5 8C9.5 8.82843 8.82843 9.5 8 9.5C7.17157 9.5 6.5 8.82843 6.5 8Z"
          ></path>
          <path
            class="dropbtn"
            d="M13 6.5C12.1716 6.5 11.5 7.17157 11.5 8C11.5 8.82843 12.1716 9.5 13 9.5C13.8284 9.5 14.5 8.82843 14.5 8C14.5 7.17157 13.8284 6.5 13 6.5Z"
          ></path>
        </svg>
        <div class="dropdown list-dropdown" style="right: 0">
          <div
            id="monitor-list-dropdown-ckcp1htucdm4n11k36sg"
            class="dropdown-content monitors-list-dropdown-content"
          >
            <a href="#" class="rename-monitor-btn">Rename</a>
            <a
              href="/viewsets/edit?id=ckcp1htucdm4n11k36sg"
              class="edit-monitor-btn">Edit</a
            >
            <a href="#" class="delete-monitor-btn">Delete</a>
          </div>
        </div>
      </div>
    </div>

    <div class="item _monitor">
      <a href="/viewsets?id=cke6koducdm31vgs1nv0">
        <div class="title">Debt</div>
        <div class="type">Viewset</div>
        <div class="visibility">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="13.5px"
            height="13.5px"
            viewBox="0 0 24 24"
            fill="none"
            style="left: -1px"
          >
            <path
              d="M7 10.0288C7.47142 10 8.05259 10 8.8 10H15.2C15.9474 10 16.5286 10 17 10.0288M7 10.0288C6.41168 10.0647 5.99429 10.1455 5.63803 10.327C5.07354 10.6146 4.6146 11.0735 4.32698 11.638C4 12.2798 4 13.1198 4 14.8V16.2C4 17.8802 4 18.7202 4.32698 19.362C4.6146 19.9265 5.07354 20.3854 5.63803 20.673C6.27976 21 7.11984 21 8.8 21H15.2C16.8802 21 17.7202 21 18.362 20.673C18.9265 20.3854 19.3854 19.9265 19.673 19.362C20 18.7202 20 17.8802 20 16.2V14.8C20 13.1198 20 12.2798 19.673 11.638C19.3854 11.0735 18.9265 10.6146 18.362 10.327C18.0057 10.1455 17.5883 10.0647 17 10.0288M7 10.0288V8C7 5.23858 9.23858 3 12 3C14.7614 3 17 5.23858 17 8V10.0288"
              stroke-width="1.8"
              stroke-linecap="round"
              stroke-linejoin="round"
            ></path>
          </svg>
          private
        </div>
        <div class="created">7 months ago</div>
      </a>
      <div onclick="dropdown(this)" class="actions dropbtn">
        <svg
          class="dropbtn"
          width="16"
          height="16"
          viewBox="0 0 16 16"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            class="dropbtn"
            d="M3 6.5C2.17157 6.5 1.5 7.17157 1.5 8C1.5 8.82843 2.17157 9.5 3 9.5C3.82843 9.5 4.5 8.82843 4.5 8C4.5 7.17157 3.82843 6.5 3 6.5Z"
          ></path>
          <path
            class="dropbtn"
            d="M6.5 8C6.5 7.17157 7.17157 6.5 8 6.5C8.82843 6.5 9.5 7.17157 9.5 8C9.5 8.82843 8.82843 9.5 8 9.5C7.17157 9.5 6.5 8.82843 6.5 8Z"
          ></path>
          <path
            class="dropbtn"
            d="M13 6.5C12.1716 6.5 11.5 7.17157 11.5 8C11.5 8.82843 12.1716 9.5 13 9.5C13.8284 9.5 14.5 8.82843 14.5 8C14.5 7.17157 13.8284 6.5 13 6.5Z"
          ></path>
        </svg>
        <div class="dropdown list-dropdown" style="right: 0">
          <div
            id="monitor-list-dropdown-cke6koducdm31vgs1nv0"
            class="dropdown-content monitors-list-dropdown-content"
          >
            <a href="#" class="rename-monitor-btn">Rename</a>
            <a
              href="/viewsets/edit?id=cke6koducdm31vgs1nv0"
              class="edit-monitor-btn">Edit</a
            >
            <a href="#" class="delete-monitor-btn">Delete</a>
          </div>
        </div>
      </div>
    </div>

    <div class="item _monitor">
      <a href="/viewsets?id=cke92ttucdmeoi777lag">
        <div class="title">Inventory</div>
        <div class="type">Viewset</div>
        <div class="visibility">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="13.5px"
            height="13.5px"
            viewBox="0 0 24 24"
            fill="none"
            style="left: -1px"
          >
            <path
              d="M7 10.0288C7.47142 10 8.05259 10 8.8 10H15.2C15.9474 10 16.5286 10 17 10.0288M7 10.0288C6.41168 10.0647 5.99429 10.1455 5.63803 10.327C5.07354 10.6146 4.6146 11.0735 4.32698 11.638C4 12.2798 4 13.1198 4 14.8V16.2C4 17.8802 4 18.7202 4.32698 19.362C4.6146 19.9265 5.07354 20.3854 5.63803 20.673C6.27976 21 7.11984 21 8.8 21H15.2C16.8802 21 17.7202 21 18.362 20.673C18.9265 20.3854 19.3854 19.9265 19.673 19.362C20 18.7202 20 17.8802 20 16.2V14.8C20 13.1198 20 12.2798 19.673 11.638C19.3854 11.0735 18.9265 10.6146 18.362 10.327C18.0057 10.1455 17.5883 10.0647 17 10.0288M7 10.0288V8C7 5.23858 9.23858 3 12 3C14.7614 3 17 5.23858 17 8V10.0288"
              stroke-width="1.8"
              stroke-linecap="round"
              stroke-linejoin="round"
            ></path>
          </svg>
          private
        </div>
        <div class="created">7 months ago</div>
      </a>
      <div onclick="dropdown(this)" class="actions dropbtn">
        <svg
          class="dropbtn"
          width="16"
          height="16"
          viewBox="0 0 16 16"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            class="dropbtn"
            d="M3 6.5C2.17157 6.5 1.5 7.17157 1.5 8C1.5 8.82843 2.17157 9.5 3 9.5C3.82843 9.5 4.5 8.82843 4.5 8C4.5 7.17157 3.82843 6.5 3 6.5Z"
          ></path>
          <path
            class="dropbtn"
            d="M6.5 8C6.5 7.17157 7.17157 6.5 8 6.5C8.82843 6.5 9.5 7.17157 9.5 8C9.5 8.82843 8.82843 9.5 8 9.5C7.17157 9.5 6.5 8.82843 6.5 8Z"
          ></path>
          <path
            class="dropbtn"
            d="M13 6.5C12.1716 6.5 11.5 7.17157 11.5 8C11.5 8.82843 12.1716 9.5 13 9.5C13.8284 9.5 14.5 8.82843 14.5 8C14.5 7.17157 13.8284 6.5 13 6.5Z"
          ></path>
        </svg>
        <div class="dropdown list-dropdown" style="right: 0">
          <div
            id="monitor-list-dropdown-cke92ttucdmeoi777lag"
            class="dropdown-content monitors-list-dropdown-content"
          >
            <a href="#" class="rename-monitor-btn">Rename</a>
            <a
              href="/viewsets/edit?id=cke92ttucdmeoi777lag"
              class="edit-monitor-btn">Edit</a
            >
            <a href="#" class="delete-monitor-btn">Delete</a>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="count">COUNT <span class="num">4</span></div>

  <div class="pagination">
    <div class="previous">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M15.75 19.5 8.25 12l7.5-7.5"
        ></path>
      </svg>
    </div>
    <div class="page active">1</div>
    <div class="page">2</div>
    <div class="page">3</div>
    <div class="page">4</div>
    <div class="next">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="m8.25 4.5 7.5 7.5-7.5 7.5"
        ></path>
      </svg>
    </div>
  </div>
</dashboards>

<style>
  @import "page.css";
</style>
