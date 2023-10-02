// CardList.js
import React, { useState, useEffect } from 'react';
import Card from './Card';

function CardList() {
  const [cards, setCards] = useState([]);

  useEffect(() => {
    async function fetchCards() {
      try {
        const response = await fetch(
          'https://bynogame-price-analyst-core.fly.dev/api/cs?limit=50'
        );
        const data = await response.json();
        setCards(data);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }

    fetchCards();
  }, []);

  return (
    <div className="card-list">
      {cards.map((card, index) => (
        <Card key={index} cardData={card} />
      ))}
    </div>
  );
}

export default CardList;
