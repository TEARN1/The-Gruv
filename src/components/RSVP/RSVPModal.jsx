import React from 'react';

export default function RSVPModal({ open, onClose, onConfirm, eventTitle }) {
  if (!open) return null;
  return (
    <div className="modal" role="dialog" aria-modal="true">
      <div className="card">
        <h3>RSVP to {eventTitle}</h3>
        <p>Confirm your RSVP. This is a lightweight modal placeholder.</p>
        <div className="row">
          <button onClick={() => { onConfirm && onConfirm(); onClose && onClose(); }}>Confirm</button>
          <button onClick={() => onClose && onClose()}>Cancel</button>
        </div>
      </div>
    </div>
  );
}