import React, { useState, useEffect } from 'react';
import axios from 'axios';
import '../styles/ProductPage.css';

const ProductPage = () => {
  const [quantity, setQuantity] = useState('');
  const [packs, setPacks] = useState([]);

  const handleChange = (e) => {
    // Only allow numbers in the input field
    setQuantity(e.target.value);
  };

  const handleSubmit = () => {
    // Perform the order placement logic here
    if (!isNaN(parseFloat(quantity)) && isFinite(quantity)) {
      axios.get('https://gym-dolphin.onrender.com/order_packs?items=' + quantity)
        // axios.get('http://localhost:8080/order_packs?items=' + quantity)
        .then(response => {
          console.log(response.data.packs)
          setPacks(response.data.packs)
        })
        .catch(error => console.error('Error fetching data:', error));
    } else {
      alert('Please enter a valid quantity');
    }
  };

  const handleKeyPress = (e) => {
    if (e.key === 'Enter') {
      e.preventDefault();
      handleSubmit();
    }
  }

  return (
    <div>
      <div className="banner">
        Gym Dolphin
      </div>

      <div className="product-container">
        <div className='image'>
          <img src='/white_shirt.jpg' width={500}></img>
        </div>

        <div className='order-container'>
          <div className="quantity-container">
            <label htmlFor="quantity" className="quantity-label">Quantity:</label>
            <input
              type="text"
              id="quantity"
              value={quantity}
              onChange={handleChange}
              onKeyPress={handleKeyPress}
              className="quantity-input"
            />
            <button className='btn quantity-btn' onClick={handleSubmit}>Place Order</button>

          </div>

        </div>
      </div>

      {packs && packs.length > 0 && (
        <div className='available-packs'>
          Available Packs:
          <ul>
            {packs.map((item, index) => (
              <li key={index}>
                Size: {item.size}, Count: {item.count}
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}


export default ProductPage;
