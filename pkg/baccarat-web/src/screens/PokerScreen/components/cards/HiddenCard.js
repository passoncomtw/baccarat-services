import React from 'react';

const HiddenCard = (props) => {
  const { 
    applyFoldedClassname
  } = props;
  return(
    <div className={`playing-card cardIn robotcard${(applyFoldedClassname ? ' folded' : '')}`} 
      style={{animationDelay: `0ms`}}>
    </div>
  )
}

export default HiddenCard;