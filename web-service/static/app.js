// The Gruv - Event Management JavaScript

class EventManager {
    constructor() {
        this.apiBaseUrl = 'http://localhost:8080/api/events';
        this.init();
    }

    init() {
        this.bindEventListeners();
        this.loadEvents();
    }

    bindEventListeners() {
        // Add Event Button
        const addEventBtn = document.getElementById('addEventBtn');
        const modal = document.getElementById('addEventModal');
        const closeBtn = modal.querySelector('.close');
        const cancelBtn = document.getElementById('cancelBtn');
        const eventForm = document.getElementById('eventForm');
        const locationTypeSelect = document.getElementById('locationType');

        addEventBtn.addEventListener('click', () => this.openModal());
        closeBtn.addEventListener('click', () => this.closeModal());
        cancelBtn.addEventListener('click', () => this.closeModal());
        
        // Close modal when clicking outside
        window.addEventListener('click', (event) => {
            if (event.target === modal) {
                this.closeModal();
            }
        });

        // Form submission
        eventForm.addEventListener('submit', (e) => this.handleFormSubmit(e));

        // Location type change
        locationTypeSelect.addEventListener('change', () => this.updateLocationField());

        // Initialize location field
        this.updateLocationField();
    }

    openModal() {
        const modal = document.getElementById('addEventModal');
        modal.style.display = 'block';
        document.body.style.overflow = 'hidden';
        
        // Set default times
        const now = new Date();
        const startTime = new Date(now.getTime() + 60 * 60 * 1000); // 1 hour from now
        const endTime = new Date(now.getTime() + 2 * 60 * 60 * 1000); // 2 hours from now
        
        document.getElementById('startTime').value = this.formatDateTimeLocal(startTime);
        document.getElementById('endTime').value = this.formatDateTimeLocal(endTime);
    }

    closeModal() {
        const modal = document.getElementById('addEventModal');
        modal.style.display = 'none';
        document.body.style.overflow = 'auto';
        
        // Reset form
        document.getElementById('eventForm').reset();
        this.updateLocationField();
    }

    updateLocationField() {
        const locationType = document.getElementById('locationType').value;
        const locationInput = document.getElementById('locationInput');
        const locationLabel = document.getElementById('locationLabel');
        
        if (locationType === 'online') {
            locationLabel.textContent = 'Meeting URL';
            locationInput.placeholder = 'https://zoom.us/j/123456789';
            locationInput.type = 'url';
        } else {
            locationLabel.textContent = 'Address';
            locationInput.placeholder = 'Enter event address';
            locationInput.type = 'text';
        }
    }

    formatDateTimeLocal(date) {
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        const hours = String(date.getHours()).padStart(2, '0');
        const minutes = String(date.getMinutes()).padStart(2, '0');
        
        return `${year}-${month}-${day}T${hours}:${minutes}`;
    }

    async handleFormSubmit(event) {
        event.preventDefault();
        
        const formData = new FormData(event.target);
        const locationType = formData.get('location_type');
        
        const eventData = {
            title: formData.get('title'),
            description: formData.get('description'),
            start_time: formData.get('start_time'),
            end_time: formData.get('end_time'),
            location: {
                type: locationType,
                [locationType === 'online' ? 'url' : 'address']: formData.get('location')
            },
            rsvp_enabled: formData.get('rsvp_enabled') === 'on',
            image_url: formData.get('image_url'),
            category: formData.get('category')
        };

        try {
            await this.createEvent(eventData);
            this.showSuccess('Event created successfully!');
            this.closeModal();
            this.loadEvents();
        } catch (error) {
            this.showError('Failed to create event: ' + error.message);
        }
    }

    async createEvent(eventData) {
        const response = await fetch(`${this.apiBaseUrl}/events`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(eventData)
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to create event');
        }

        return response.json();
    }

    async loadEvents() {
        try {
            this.showLoading();
            const response = await fetch(`${this.apiBaseUrl}/events`);
            
            if (!response.ok) {
                throw new Error('Failed to fetch events');
            }
            
            const data = await response.json();
            this.displayEvents(data.events || []);
        } catch (error) {
            console.error('Error loading events:', error);
            this.showError('Failed to load events: ' + error.message);
        }
    }

    displayEvents(events) {
        const eventsList = document.getElementById('eventsList');
        
        if (!events || events.length === 0) {
            eventsList.innerHTML = `
                <div class="no-events">
                    <i class="fas fa-calendar-alt" style="font-size: 3rem; color: #ccc; margin-bottom: 1rem;"></i>
                    <p>No events yet. Create your first event!</p>
                </div>
            `;
            return;
        }

        eventsList.innerHTML = events.map(event => this.createEventCard(event)).join('');
    }

    createEventCard(event) {
        const startDate = new Date(event.start_time);
        const endDate = new Date(event.end_time);
        const locationText = event.location.type === 'online' ? 
            `🌐 Online Event` : 
            `📍 ${event.location.address || 'Physical Location'}`;

        return `
            <div class="event-card">
                ${event.image_url ? `<img src="${event.image_url}" alt="${event.title}" style="width: 100%; height: 200px; object-fit: cover; border-radius: 10px; margin-bottom: 1rem;">` : ''}
                <div class="event-header">
                    <h3 style="margin-bottom: 0.5rem; color: #333;">${event.title}</h3>
                    <span class="event-category" style="background: #667eea; color: white; padding: 0.25rem 0.5rem; border-radius: 15px; font-size: 0.8rem; text-transform: capitalize;">
                        ${event.category}
                    </span>
                </div>
                <p style="color: #666; margin-bottom: 1rem; line-height: 1.5;">${event.description || 'No description provided.'}</p>
                <div class="event-details" style="font-size: 0.9rem; color: #555; line-height: 1.6;">
                    <div style="margin-bottom: 0.5rem;">
                        <i class="fas fa-calendar"></i> 
                        ${startDate.toLocaleDateString()} ${startDate.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'})} - 
                        ${endDate.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'})}
                    </div>
                    <div style="margin-bottom: 0.5rem;">
                        ${locationText}
                    </div>
                    ${event.rsvp_enabled ? '<div><i class="fas fa-check-circle" style="color: #2ed573;"></i> RSVP Enabled</div>' : ''}
                </div>
            </div>
        `;
    }

    showLoading() {
        const eventsList = document.getElementById('eventsList');
        eventsList.innerHTML = `
            <div class="loading">
                <i class="fas fa-spinner fa-spin" style="font-size: 2rem; margin-bottom: 1rem;"></i>
                <p>Loading events...</p>
            </div>
        `;
    }

    showError(message) {
        this.showMessage(message, 'error');
    }

    showSuccess(message) {
        this.showMessage(message, 'success');
    }

    showMessage(message, type) {
        // Remove existing messages
        const existingMessages = document.querySelectorAll('.error, .success');
        existingMessages.forEach(msg => msg.remove());

        // Create new message
        const messageDiv = document.createElement('div');
        messageDiv.className = type;
        messageDiv.textContent = message;
        
        // Insert at the top of the events section
        const eventsSection = document.querySelector('.events-section');
        eventsSection.insertBefore(messageDiv, eventsSection.firstChild);

        // Auto-remove after 5 seconds
        setTimeout(() => {
            messageDiv.remove();
        }, 5000);
    }
}

// Initialize the app when the page loads
document.addEventListener('DOMContentLoaded', () => {
    new EventManager();
});