import React, { useState } from 'react';

export default function GruvDropClaim({ drop }) {
  const [claimed, setClaimed] = useState(false);
  const handleClaim = () => {
    // placeholder: call API to claim
    setClaimed(true);
    alert('Claimed (placeholder)');
  };

  return (
    <div className="card">
      <h4>{drop?.title || 'GruvDrop'}</h4>
      <p>{drop?.description || 'Limited-time claim.'}</p>
      <div className="row">
        <button onClick={handleClaim} disabled={claimed}>{claimed ? 'Claimed' : 'Claim'}</button>
      </div>
    </div>
  );
}