import React from "react";

interface FormInputProps {
  label: string;
  type: string;
  help?: string;
}

const FormInput: React.FC<FormInputProps> = ({ label, type, help }) => {
  return (
    <div className="flex flex-col my-1 gap-1">
        <label className="text-gray-800">
            {label} <span className="text-red-500">*</span>
        </label>
        <input
            type={type}
            className="border border-gray-300 rounded-[3px] focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <span className="text-sm text-gray-400">
            {help}
        </span>
    </div>
  );
};

export default FormInput;
