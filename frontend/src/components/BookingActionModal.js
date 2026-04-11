import React from 'react';

const BookingActionModal = ({
  show,
  actionType,
  selectedBooking,
  notes,
  onNotesChange,
  onClose,
  onSubmit,
  submitting,
  getServiceName,
  getStudentName,
}) => {
  if (!show || !selectedBooking) return null;

  const isApprove = actionType === 'approve';

  return (
    <div
      className="modal-overlay"
      data-testid="booking-action-modal-overlay"
      onClick={onClose}
    >
      <div
        className="modal"
        data-testid="booking-action-modal"
        onClick={(e) => e.stopPropagation()}
      >
        <div className="modal-header">
          <h3>{isApprove ? 'Approve Booking' : 'Reject Booking'}</h3>
          <button
            className="close-btn"
            onClick={onClose}
            disabled={submitting}
            aria-label="Close modal"
          >
            ✕
          </button>
        </div>

        <div className="modal-body">
          <p className="booking-info">
            <strong>{getServiceName(selectedBooking)}</strong> —{' '}
            {getStudentName(selectedBooking)}
          </p>

          <div className="form-group">
            <label htmlFor="approval-notes">Additional Notes (Optional)</label>
            <textarea
              id="approval-notes"
              value={notes}
              onChange={(e) => onNotesChange(e.target.value)}
              placeholder="Add any notes about this approval or rejection..."
              rows="4"
              disabled={submitting}
            />
          </div>

          <div className="modal-actions">
            <button
              className="btn-cancel"
              onClick={onClose}
              disabled={submitting}
            >
              Cancel
            </button>

            <button
              className={isApprove ? 'btn-approve' : 'btn-reject'}
              onClick={onSubmit}
              disabled={submitting}
            >
              {submitting
                ? 'Submitting...'
                : isApprove
                ? 'Approve Booking'
                : 'Reject Booking'}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default BookingActionModal;