# <div align="center">GoTweet - A Minimalist Social Media Platform</div>

[GoTweet (Deployed)](http://tweets.beneboba.me:8080) is a lightweight social media platform inspired by Twitter, designed for simplicity and efficiency. Users can read and post tweets, engage with the community, comment and rate tweets, and stay updated on the latest feeds. The platform is built using Golang with the Fiber v2 framework, Gorm for database interactions, and features JWT token-based authentication. The minimalist design ensures a seamless user experience, and HTML template engine is employed for rendering views.

## <div align="center">Features</div>

- **User Authentication:** Secure user authentication powered by JWT tokens ensures a safe and private experience.

- **Tweet Feeds:** Browse through a clean and clutter-free timeline of tweets, allowing users to stay connected with the community.

- **Tweet Creation:** Post tweets and share your thoughts with the world. The platform enables users to express themselves with ease.

- **Tweet Comments and Ratings:** Engage in conversations by commenting on tweets and providing ratings. Each tweet accumulates points based on user ratings.

- **Account Management:** Create a new account or log in to an existing one. Manage your profile and stay in control of your account.

- **Minimalist Design:** The user interface is designed to be minimalistic, providing a focused and distraction-free environment.

## <div align="center">Tech Stack</div>

- **Backend:** Golang, Fiber v2, Gorm (for database interactions)
  
- **Authentication:** JWT Token-based authentication
  
- **Frontend:** HTML template engine (for minimalist and efficient views)

## <div align="center">Getting Started</div>

### Prerequisites

Make sure you have Golang installed on your machine.

### Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/benebobaa/harisenin-mini-project
    cd harisenin-mini-project
    ```

2. **Install dependencies:**

    ```bash
    go get -u
    ```

3. **Set up the database:**

    Configure your database settings in the `config` file.

4. **Run the application:**

    ```bash
    go run main.go
    ```

   The application will be running on http://localhost:8080.

## <div align="center">Usage</div>

1. **Create a new account:**

    Access the platform, click on the sign-up option, and fill in the required details.

2. **Log in:**

    Use your credentials to log in and start exploring the tweet feeds.

3. **Post a tweet:**

    Share your thoughts by creating and posting tweets directly from the platform.

4. **Comment and Rate Tweets:**

    Engage with other users by commenting on their tweets and providing ratings. Each tweet accumulates points based on user ratings.

## <div align="center">Deployment</div>

Deploy GoTweet on your preferred server or platform to make it accessible globally. The live version can be accessed [here](http://tweets.beneboba.me:8080).


## <div align="center">Contributing</div>

Feel free to contribute to the project by following the standard contribution guidelines. Fork the repository, make your changes, and submit a pull request.

## <div align="center">License</div>

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## <div align="center">Acknowledgments</div>

- Shoutout to the Golang community
- Inspiration from Twitter
- etc.

---

Feel free to further customize this README to match your project's specific features and requirements. Consider providing more details on the comment and rating system, such as how the points are calculated, or any specific functionalities related to this feature.
