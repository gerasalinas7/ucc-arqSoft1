import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './components/Login';
import HelloWorld from './components/HelloWorld';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/hello" element={<HelloWorld />} />
      </Routes>
    </Router>
  );
}

export default App;
