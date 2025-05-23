﻿using DistComp_1.DTO.RequestDTO;
using FluentValidation;

namespace DistComp_1.Infrastructure.Validators;

public class NewsRequestDTOValidator : AbstractValidator<NewsRequestDTO>
{
    public NewsRequestDTOValidator()
    {
        RuleFor(dto => dto.Title).Length(2, 64);
        RuleFor(dto => dto.Content).Length(4, 2048);
    }
}