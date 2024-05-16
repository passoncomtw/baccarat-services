import React from 'react';

const HiddenCard = (props) => {
  const { 
    applyFoldedClassname,
    containerStyle = {},
  } = props;
  return(
    <div className={`playing-card cardIn robotcard${(applyFoldedClassname ? ' folded' : '')}`} 
      style={{animationDelay: `0ms`, ...containerStyle}}>
    </div>
  )
}

export default HiddenCard;