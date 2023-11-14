<a name="readme-top"></a>

<br />
<div align="center">
  <a href="https://github.com/Childebrand94/micro-reddit">
    <img src="./frontend/public/assets/Reddit-Logo.wine.svg" alt="Logo" width="300" height="200"> 
  </a>
<h3 align="center">Micro Reddit</h3>
  <p align="center">
    A full-stack Reddit clone focusing on core functionalities like post creation, commenting, and voting, built using Go, PostgreSQL, Vite, React, and Tailwind CSS.
    <br />
    <a href="https://github.com/Childebrand94/micro-reddit"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/Childebrand94/micro-reddit">View Demo</a>
    ·
    <a href="https://github.com/Childebrand94/micro-reddit/issues">Report Bug</a>
    ·
    <a href="https://github.com/Childebrand94/micro-reddit/issues">Request Feature</a>
  </p>
</div>
<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>
<!-- ABOUT THE PROJECT -->
About The Project

![Alt Text](frontend/public/assets/showcaseGif.gif)

Micro-reddit is my first full-stack project, a Reddit clone with core functionalities like account creation, posting, commenting, and a voting system. The backend is built with Go, using PostgreSQL and pgx, while the frontend utilizes Vite, React, and Tailwind CSS. The project aims to replicate the classic Reddit experience with a modern touch, featuring a user-friendly interface and robust performance.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

-   [![Go][Go]][Go-url]
-   [![PostgreSQL][PostgreSQL]][PostgreSQL-url]
-   [![Vite][Vite]][Vite-url]
-   [![React][React.js]][React-url]
-   [![TypeScript][TypeScript]][TypeScript-url]
-   [![TailwindCSS][TailwindCSS]][TailwindCSS-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

Welcome to Micro-Reddit! This project is a compact version of Reddit, showcasing the integration of modern web technologies. The front-end is built with Vite, React, TypeScript, and Tailwind CSS, offering a responsive and dynamic user interface. The back-end is powered by Go with a PostgreSQL database, ensuring efficient data handling and scalability. Here's how to set up and run the project on your local machine for development and testing.

### Prerequisites

Before you begin, ensure you have the following prerequisites installed:

-   [Node.js](https://nodejs.org/en/download/)
-   [Go](https://go.dev/dl/)
-   [PostgreSQL](https://www.postgresql.org/)
-   [NPM](https://www.npmjs.com/)

### Installation

1. **Clone the Repo**
   Clone the repository to your local machine:
    ```sh
    git clone https://github.com/Childebrand94/micro-reddit
    ```
2. Front-End Setup:
   Navigate to the front-end directory and install the dependencies:
    ```sh
    cd frontend
    npm install
    ```
3. Run the file
    ```sh
    npm run dev
    ```
4. Back-End Setup:
   In a separate terminal, navigate to the back-end directory. Install Go dependencies:
    ```sh
    cd backend
    go mod tidy
    ```
5. Database Setup:
   Ensure PostgreSQL is running. To create schema and populate with fake data.
    ```sh
    cd backend
    make migrate
    make seed
    ```
6. Front-End: Run the Vite development server:
    ```sh
    cd frontend
    npm run dev
    ```
7. Back-end: Start up server after schema has be created and populated
    ```sh
    cd backend
    make run
    ```
8. Go to http://localhost:5173/

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

Micro-reddit offers an intuitive and engaging platform for users to interact with a community-driven content system, similar to the classic Reddit experience. Once registered, users can create posts, comment on existing posts, and participate in discussions by upvoting or downvoting content. The application features a straightforward navigation system, allowing users to easily browse through different posts, utilize the search bar for specific content, and engage with the community. Whether you're looking to share interesting links, start discussions, or simply explore community content, micro-reddit provides a seamless and user-friendly environment. The sorting system ensures that users can access the most relevant and up-to-date content, enhancing the overall user experience on the platform.

<!-- ROADMAP -->

## Roadmap

-   [x] Arrow's highlight showing users pervious vote
-   [x] Hierarchical comments
-   [ ] Expand/Collapse comment replies
-   [ ] Pagination

See the [open issues](https://github.com/Childebrand94/micro-reddit/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Feedback and Contributions

We value your feedback and encourage you to report any issues, suggest features, or contribute to the development of this project. Together, we can make this Blackjack Trainer the ultimate resource for mastering the game.

Please note that this roadmap is subject to change, and the timeline for each feature may vary. We appreciate your support and look forward to improving and expanding our Blackjack Trainer based on your needs and suggestions.

See the [open issues](https://github.com/Childebrand94/micro-reddit/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->

[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[NPM]: https://img.shields.io/badge/npm-CB3837?style=for-the-badge&logo=npm&logoColor=white
[NPM-url]: https://www.npmjs.com/
[TailwindCSS]: https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white
[TailwindCSS-url]: https://v2.tailwindcss.com/docs
[Vite]: https://img.shields.io/badge/vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white
[Vite-url]: https://vitejs.dev/
[Go]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
[ReactRouter]: https://img.shields.io/badge/React_Router-CA4245?style=for-the-badge&logo=react-router&logoColor=white
[ReactRouter-url]: https://reactrouter.com/en/main
[PostgreSQL]: https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white
[PostgreSQL-url]: https://www.postgresql.org/
[TypeScript]: https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white
[TypeScript-url]: https://www.typescriptlang.org/
