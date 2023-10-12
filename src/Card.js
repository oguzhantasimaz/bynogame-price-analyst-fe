// Card.js
import React, { useState } from 'react';

function Card({ cardData }) {
  const [imageError, setImageError] = useState(false);

  const handleImageError = () => {
    setImageError(true);
  };

  const queryEscape = (string) => {
    var params = new URLSearchParams();
    params.append('q', string);

    let str = params.toString();
    str = string.replace("q=", "")

    return str
  
  }


  // Calculate the percentage difference between 'price' and 'steamStats.lowest_price'
  const calculatePriceDifference = () => {
    const originalPrice = cardData.price;
    const lowestPrice = parseFloat(
      cardData.steamStats.lowest_price.replace('.', ',').replace(' TL', '')
    );
    const percentageDifference = ((lowestPrice - originalPrice) / originalPrice) * 100;
    return percentageDifference.toFixed(2);
  };

  const isSold = cardData.dateTimeSold > 0;

  return (
    <div className={`card ${isSold ? 'sold' : ''}`}>
      {isSold && (
        <div className="sold-flag">SOLD!</div>
      )}
      {imageError ? (
        <img
          src="path_to_placeholder_image.png" // Provide the path to your placeholder image
          alt={cardData.nameSlug}
        />
      ) : (
        <img
          src={cardData.image}
          alt={cardData.nameSlug}
          onError={handleImageError}
        />
      )}
      <h2>{cardData.typeInfoSteam.hash}</h2>
      <p>Description: {cardData.description}</p>
      <p>Float: {cardData.float || "Unavailable"}</p>
      <p>ByNoGame Price: {cardData.price} TL</p>
      <p>Steam Latest Price: {cardData.steamStats.lowest_price}</p>
      <p>Price Difference: <b>{calculatePriceDifference()}%</b></p>
      <p>Steam Daily Volume: {cardData.steamStats.volume}</p>
      <p>Steam Median Price: {cardData.steamStats.median_price}</p>
      <p>Seller Name: {cardData.sellerMarketName}</p>
      <a href={"https://www.bynogame.com/en/games/csgo/skin/" + cardData.nameSlug} target="_blank" rel="noopener noreferrer">Buy From ByNoGame</a>
      <br />
      <br />
      <a href={"https://steamcommunity.com/market/listings/730/" + queryEscape(cardData.typeInfoSteam.hash)} target="_blank" rel="noopener noreferrer">Buy From Steam</a>
      {/* Add more card information as needed */}
    </div>
  );
}

export default Card;