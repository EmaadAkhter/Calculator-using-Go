# Calculator using Go 🧮

A beautiful, responsive calculator web application built with Go backend and vanilla HTML/CSS/JavaScript frontend. This project demonstrates full-stack development with a REST API backend and an interactive web interface.

## 🚀 Live Demo

**Website:** [https://calculator-bbg3.onrender.com/](https://calculator-bbg3.onrender.com/)

## ✨ Features

- **Beautiful UI**: Modern, colorful design with smooth animations and hover effects
- **Responsive Design**: Works perfectly on desktop, tablet, and mobile devices
- **Real-time Calculations**: Performs calculations via REST API calls to Go backend
- **Keyboard Support**: Full keyboard navigation and input support
- **Error Handling**: Comprehensive error messages for invalid operations
- **CORS Enabled**: Frontend can communicate seamlessly with backend API
- **Loading States**: Visual feedback during API calls

## 🛠️ Tech Stack

### Backend
- **Go 1.21**: Main programming language
- **net/http**: Built-in HTTP server and routing
- **encoding/json**: JSON request/response handling
- **CORS**: Cross-origin resource sharing support

### Frontend
- **HTML5**: Semantic markup
- **CSS3**: Modern styling with gradients, animations, and flexbox/grid
- **Vanilla JavaScript**: Interactive functionality and API communication
- **Google Fonts**: Quicksand font family for beautiful typography

### Deployment
- **Render**: Cloud platform for both backend API and static frontend
- **GitHub**: Version control and CI/CD integration

## 📁 Project Structure

```
calculator-using-go/
├── main.go              # Go backend server
├── go.mod              # Go module file
├── static/
│   └── index.html      # Frontend application
├── .gitignore          # Git ignore rules
├── LICENSE             # MIT License
└── README.md           # Project documentation
```

## 🔧 API Endpoints

### Base URL: `https://calculator-using-go.onrender.com`

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Welcome message |
| `/hello` | GET | Simple hello world |
| `/calc` | POST | Perform calculation |

### Calculator API (`/calc`)

**Request Format:**
```json
{
    "num1": 10,
    "num2": 5,
    "oprator": "+"
}
```

**Supported Operations:**
- `+` : Addition
- `-` : Subtraction  
- `*` : Multiplication
- `/` : Division

**Response Format:**
```
The Result of Addition is:15
```

**Error Handling:**
- `400 Bad Request`: Invalid JSON, missing parameters, division by zero, invalid operation
- `200 OK`: Successful calculation

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher
- Git

### Local Development

1. **Clone the repository:**
```bash
git clone https://github.com/Emaad_Ansari/calculator-using-go.git
cd calculator-using-go
```

2. **Run the backend:**
```bash
go mod tidy
go run main.go
```

3. **Access the application:**
- Backend API: `http://localhost:8080`
- Frontend: Open `static/index.html` in your browser
- For local development, update the API URL in `index.html` to `http://localhost:8080`

### Testing the API

**Using curl:**
```bash
# Addition
curl -X POST http://localhost:8080/calc \
  -H "Content-Type: application/json" \
  -d '{"num1": 10, "num2": 5, "oprator": "+"}'

# Division
curl -X POST http://localhost:8080/calc \
  -H "Content-Type: application/json" \
  -d '{"num1": 20, "num2": 4, "oprator": "/"}'
```

## 🎨 UI Features

- **Glassmorphism Design**: Modern translucent design with backdrop blur effects
- **Smooth Animations**: Hover effects, button press animations, and loading states
- **Color Palette**: Soft pastels including pink, blue, and purple gradients
- **Interactive Elements**: Responsive buttons with visual feedback
- **Error States**: Clear error messaging with distinct styling
- **Accessibility**: Proper contrast ratios and keyboard navigation

## 🔒 Security Features

- **CORS Headers**: Properly configured for cross-origin requests
- **Input Validation**: Server-side validation for all calculator inputs
- **Error Handling**: Graceful error responses without exposing internal details
- **Division by Zero Protection**: Prevents mathematical errors

## 🚀 Deployment

This project is deployed on Render:

- **Backend**: Deployed as a Go web service
- **Frontend**: Deployed as a static site
- **Auto-deployment**: Triggered on GitHub commits to main branch

### Deploy Your Own

1. **Fork this repository**
2. **Backend Deployment:**
   - Create new Web Service on Render
   - Connect your GitHub repo
   - Use Go environment
   - Build command: `go build -o main`
   - Start command: `./main`

3. **Frontend Deployment:**
   - Create new Static Site on Render
   - Connect your GitHub repo
   - Publish directory: `static`
   - No build command needed

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🐛 Known Issues

- Calculator currently works with integers only
- No decimal point calculations in backend (frontend shows decimal button)
- No memory functions (M+, M-, MR, MC)

## 🔮 Future Enhancements

- [ ] Support for decimal numbers
- [ ] Memory functions
- [ ] Calculation history
- [ ] Scientific calculator mode
- [ ] Dark/light theme toggle
- [ ] Expression evaluation (e.g., "2+3*4")
- [ ] Unit tests for backend
- [ ] Docker containerization

## 👨‍💻 Author

**Emaad Ansari**
- GitHub: [@Emaad_Ansari](https://github.com/Emaad_Ansari)

## ⭐ Show Your Support

Give a ⭐ if this project helped you learn Go web development!

---

**Built with ❤️ using Go and modern web technologies**
