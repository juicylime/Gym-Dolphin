import './App.css';
import ProductPage from './components/ProductPage.js';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/" element={<ProductPage />} />
        </Routes>
      </Router>
    </div>
    
  );
}

export default App;
