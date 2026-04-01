'use client';
import React, { useState } from 'react';
import SelectService from './selectservice';
import SelectedTimeService from './selectedtimeservice';
import ServiceSidebar from '../../components/patientcomponent/servicesideber';

export default function BookingPage() {
  const [step, setStep] = useState(1);
  const [selectedService, setSelectedService] = useState('');
  const [selectedDate, setSelectedDate] = useState(4);
  const [selectedTime, setSelectedTime] = useState('');

  return (
    <div className="min-h-screen bg-[#eeeeee] flex flex-col md:flex-row">
      {step === 1 ? (
        <SelectService selected={selectedService} onSelect={setSelectedService} />
      ) : (
        <SelectedTimeService 
          date={selectedDate} onDateChange={setSelectedDate}
          time={selectedTime} onTimeChange={setSelectedTime}
        />
      )}
      
      <ServiceSidebar 
        service={selectedService} 
        date={step === 2 ? selectedDate : null}
        time={selectedTime}
        onNext={() => step === 1 ? setStep(2) : alert('ยืนยันการจอง')}
      />
    </div>
  );
}