package w32

// WINLANG_H
const (
	LANG_NEUTRAL        = 0x00 // Default custom (MUI) locale language
	LANG_USER_DEFAULT   = 0x01 // User default locale language
	LANG_SYSTEM_DEFAULT = 0x02 // System default locale language
	LANG_INVARIANT      = 0x7F // Invariant locale language

	SUBLANG_NEUTRAL            = 0x00 // Neutral sublanguage
	SUBLANG_INVARIANT          = 0x00 // Invariant sublanguage
	SUBLANG_DEFAULT            = 0x01 // User default sublanguage
	SUBLANG_SYS_DEFAULT        = 0x02 // System default sublanguage
	SUBLANG_CUSTOM_DEFAULT     = 0x03 // Default custom sublanguage
	SUBLANG_CUSTOM_UNSPECIFIED = 0x04 // Unspecified custom sublanguage
	SUBLANG_UI_CUSTOM_DEFAULT  = 0x05 // Default custom MUI sublanguage

	/*
	 * All (sub)language identifiers
	 */
	LANG_AFRIKAANS                 = 0x36 // Afrikaans (af)
	SUBLANG_AFRIKAANS_SOUTH_AFRICA = 0x01 // South Africa (ZA)

	LANG_ALBANIAN            = 0x1C // Albanian (sq)
	SUBLANG_ALBANIAN_ALBANIA = 0x01 // Albania (AL)

	LANG_ALSATIAN           = 0x84 // Alsatian (gsw)
	SUBLANG_ALSATIAN_FRANCE = 0x01 // France (FR)

	LANG_AMHARIC             = 0x5E // Amharic (am)
	SUBLANG_AMHARIC_ETHIOPIA = 0x01 // Ethiopia (ET)

	LANG_ARABIC                 = 0x01 // Arabic (ar)
	SUBLANG_ARABIC_ALGERIA      = 0x05 // Algeria (DZ)
	SUBLANG_ARABIC_BAHRAIN      = 0x0F // Bahrain (BH)
	SUBLANG_ARABIC_EGYPT        = 0x03 // Egypt (EG)
	SUBLANG_ARABIC_IRAQ         = 0x02 // Iraq (IQ)
	SUBLANG_ARABIC_JORDAN       = 0x0B // Jordan (JO)
	SUBLANG_ARABIC_KUWAIT       = 0x0D // Kuwait (KW)
	SUBLANG_ARABIC_LEBANON      = 0x0C // Lebanon (LB)
	SUBLANG_ARABIC_LIBYA        = 0x04 // Libya (LY)
	SUBLANG_ARABIC_MOROCCO      = 0x06 // Morocco (MA)
	SUBLANG_ARABIC_OMAN         = 0x08 // Oman (OM)
	SUBLANG_ARABIC_QATAR        = 0x10 // Qatar (QA)
	SUBLANG_ARABIC_SAUDI_ARABIA = 0x01 // Saudi Arabia (SA)
	SUBLANG_ARABIC_SYRIA        = 0x0A // Syria (SY)
	SUBLANG_ARABIC_TUNISIA      = 0x07 // Tunisia (TN)
	SUBLANG_ARABIC_UAE          = 0x0E // U.A.E. (AE)
	SUBLANG_ARABIC_YEMEN        = 0x09 // Yemen (YE)

	LANG_ARMENIAN            = 0x2B // Armenian (hy)
	SUBLANG_ARMENIAN_ARMENIA = 0x01 // Armenia (AM)

	LANG_ASSAMESE          = 0x4D // Assamese (as)
	SUBLANG_ASSAMESE_INDIA = 0x01 // India (IN)

	LANG_AZERI             = 0x2C // Azerbaijani (az)
	SUBLANG_AZERI_CYRILLIC = 0x02 // Azerbaijan, Cyrillic (AZ)
	SUBLANG_AZERI_LATIN    = 0x01 // Azerbaijan, Latin (AZ)

	LANG_BANGLA               = 0x45 // Bangla (bn)
	SUBLANG_BANGLA_BANGLADESH = 0x02 // Bangladesh
	SUBLANG_BANGLA_INDIA      = 0x01 // India (IN)

	LANG_BASHKIR           = 0x6D // Bashkir (ba)
	SUBLANG_BASHKIR_RUSSIA = 0x01 // Russia (RU)

	LANG_BASQUE           = 0x2D // Basque (Basque)
	SUBLANG_BASQUE_BASQUE = 0x01 // Basque (Basque)

	LANG_BELARUSIAN            = 0x23 // Belarusian (be)
	SUBLANG_BELARUSIAN_BELARUS = 0x01 // Belarus (BY)

	LANG_BOSNIAN_NEUTRAL                        = 0x781A // Bosnian (bs) - Neutral
	LANG_BOSNIAN                                = 0x1A   // Bosnian (bs)
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_CYRILLIC = 0x08   // Bosnia and Herzegovina, Cyrillic (BA)
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_LATIN    = 0x05   // Bosnia and Herzegovina, Latin (BA)

	LANG_BRETON           = 0x7E // Breton (br)
	SUBLANG_BRETON_FRANCE = 0x01 // France (FR)

	LANG_BULGARIAN             = 0x02 // Bulgarian (bg)
	SUBLANG_BULGARIAN_BULGARIA = 0x01 // Bulgaria (BG)

	LANG_CENTRAL_KURDISH         = 0x92 // Central Kurdish (ku)
	SUBLANG_CENTRAL_KURDISH_IRAQ = 0x01 // Iraq (IQ)

	LANG_CHEROKEE             = 0x5C // Cherokee (chr)
	SUBLANG_CHEROKEE_CHEROKEE = 0x01 // Cherokee (Cher)

	LANG_CATALAN            = 0x03 // Catalan (ca)
	SUBLANG_CATALAN_CATALAN = 0x01 // Spain (ES)

	LANG_CHINESE              = 0x04 // Chinese (zh)
	SUBLANG_CHINESE_HONGKONG  = 0x03 // Hong Kong SAR, PRC (HK)
	SUBLANG_CHINESE_MACAU     = 0x05 // Macao SAR (MO)
	SUBLANG_CHINESE_SINGAPORE = 0x04 // Singapore (SG)

	LANG_CHINESE_SIMPLIFIED    = 0x04 // Chinese (zh)
	SUBLANG_CHINESE_SIMPLIFIED = 0x02 // Simplified (Hans)

	LANG_CHINESE_TRADITIONAL    = 0x7C04 // Chinese (zh)
	SUBLANG_CHINESE_TRADITIONAL = 0x7C04 // Traditional (Hant)

	LANG_CORSICAN           = 0x83 // Corsican (co)
	SUBLANG_CORSICAN_FRANCE = 0x01 // France (FR)

	LANG_CROATIAN                             = 0x1A // Croatian (hr)
	SUBLANG_CROATIAN_BOSNIA_HERZEGOVINA_LATIN = 0x04 // Bosnia and Herzegovina, Latin (BA)
	SUBLANG_CROATIAN_CROATIA                  = 0x01 // Croatia (HR)

	LANG_CZECH                   = 0x05 // Czech (cs)
	SUBLANG_CZECH_CZECH_REPUBLIC = 0x01 // Czech Republic (CZ)

	LANG_DANISH            = 0x06 // Danish (da)
	SUBLANG_DANISH_DENMARK = 0x01 // Denmark (DK)

	LANG_DARI                = 0x8C // Dari (prs)
	SUBLANG_DARI_AFGHANISTAN = 0x01 // Afghanistan (AF)

	LANG_DIVEHI             = 0x65 // Divehi (dv)
	SUBLANG_DIVEHI_MALDIVES = 0x01 // Maldives (MV)

	LANG_DUTCH            = 0x13 // Dutch (nl)
	SUBLANG_DUTCH_BELGIAN = 0x02 // Belgium (BE)
	SUBLANG_DUTCH         = 0x01 // Netherlands (NL)

	LANG_ENGLISH                 = 0x09 // English (en)
	SUBLANG_ENGLISH_AUS          = 0x03 // Australia (AU)
	SUBLANG_ENGLISH_BELIZE       = 0x0A // Belize (BZ)
	SUBLANG_ENGLISH_CAN          = 0x04 // Canada (CA)
	SUBLANG_ENGLISH_CARIBBEAN    = 0x09 // Caribbean (029)
	SUBLANG_ENGLISH_INDIA        = 0x10 // India (IN)
	SUBLANG_ENGLISH_EIRE         = 0x06 // Ireland (IE)
	SUBLANG_ENGLISH_IRELAND      = 0x06 // Ireland (IE)
	SUBLANG_ENGLISH_JAMAICA      = 0x08 // Jamaica (JM)
	SUBLANG_ENGLISH_MALAYSIA     = 0x11 // Malaysia (MY)
	SUBLANG_ENGLISH_NZ           = 0x05 // New Zealand (NZ)
	SUBLANG_ENGLISH_PHILIPPINES  = 0x0D // Philippines (PH)
	SUBLANG_ENGLISH_SINGAPORE    = 0x12 // Singapore (SG)
	SUBLANG_ENGLISH_SOUTH_AFRICA = 0x07 // South Africa (ZA)
	SUBLANG_ENGLISH_TRINIDAD     = 0x0B // Trinidad and Tobago (TT)
	SUBLANG_ENGLISH_UK           = 0x02 // United Kingdom (GB)
	SUBLANG_ENGLISH_US           = 0x01 // United States (US)
	SUBLANG_ENGLISH_ZIMBABWE     = 0x0C // Zimbabwe (ZW)

	LANG_ESTONIAN            = 0x25 // Estonian (et)
	SUBLANG_ESTONIAN_ESTONIA = 0x01 // Estonia (EE)

	LANG_FAEROESE                  = 0x38 // Faroese (fo)
	SUBLANG_FAEROESE_FAROE_ISLANDS = 0x01 // Faroe Islands (FO)

	LANG_FILIPINO                = 0x64 // Filipino (fil)
	SUBLANG_FILIPINO_PHILIPPINES = 0x01 // Philippines (PH)

	LANG_FINNISH            = 0x0B // Finnish (fi)
	SUBLANG_FINNISH_FINLAND = 0x01 // Finland (FI)

	LANG_FRENCH               = 0x0C // French (fr)
	SUBLANG_FRENCH_BELGIAN    = 0x02 // Belgium (BE)
	SUBLANG_FRENCH_CANADIAN   = 0x03 // Canada (CA)
	SUBLANG_FRENCH            = 0x01 // France (FR)
	SUBLANG_FRENCH_LUXEMBOURG = 0x05 // Luxembourg (LU)
	SUBLANG_FRENCH_MONACO     = 0x06 // Monaco (MC)
	SUBLANG_FRENCH_SWISS      = 0x04 // Switzerland (CH)

	LANG_FRISIAN                = 0x62 // Frisian (fy)
	SUBLANG_FRISIAN_NETHERLANDS = 0x01 // Netherlands (NL)

	LANG_GALICIAN             = 0x56 // Galician (gl)
	SUBLANG_GALICIAN_GALICIAN = 0x01 // Spain (ES)

	LANG_GEORGIAN            = 0x37 // Georgian (ka)
	SUBLANG_GEORGIAN_GEORGIA = 0x01 // Georgia (GE)

	LANG_GERMAN                  = 0x07 // German (de)
	SUBLANG_GERMAN               = 0x01 // Germany (DE)
	SUBLANG_GERMAN_SWISS         = 0x02 // Switzerland (CH)
	SUBLANG_GERMAN_AUSTRIAN      = 0x03 // Austria (AT)
	SUBLANG_GERMAN_LUXEMBOURG    = 0x04 // Luxembourg (LU)
	SUBLANG_GERMAN_LIECHTENSTEIN = 0x05 // Liechtenstein (LI)

	LANG_GREEK           = 0x08 // Greek (el)
	SUBLANG_GREEK_GREECE = 0x01 // Greece (GR)

	LANG_GREENLANDIC              = 0x6F // Greenlandic (kl)
	SUBLANG_GREENLANDIC_GREENLAND = 0x01 // Greenland (GL)

	LANG_GUJARATI          = 0x47 // Gujarati (gu)
	SUBLANG_GUJARATI_INDIA = 0x01 // India (IN)

	LANG_HAUSA                  = 0x68 // Hausa (ha)
	SUBLANG_HAUSA_NIGERIA_LATIN = 0x01 // Nigeria (NG)

	LANG_HAWAIIAN       = 0x75 // Hawiian (haw)
	SUBLANG_HAWAIIAN_US = 0x01 // United States (US)

	LANG_HEBREW           = 0x0D // Hebrew (he)
	SUBLANG_HEBREW_ISRAEL = 0x01 // Israel (IL)

	LANG_HINDI          = 0x39 // Hindi (hi)
	SUBLANG_HINDI_INDIA = 0x01 // India (IN)

	LANG_HUNGARIAN            = 0x0E // Hungarian (hu)
	SUBLANG_HUNGARIAN_HUNGARY = 0x01 // Hungary (HU)

	LANG_ICELANDIC            = 0x0F // Icelandic (is)
	SUBLANG_ICELANDIC_ICELAND = 0x01 // Iceland (IS)

	LANG_IGBO            = 0x70 // Igbo (ig)
	SUBLANG_IGBO_NIGERIA = 0x01 // Nigeria (NG)

	LANG_INDONESIAN              = 0x21 // Indonesian (id)
	SUBLANG_INDONESIAN_INDONESIA = 0x01 // Indonesia (ID)

	LANG_INUKTITUT                 = 0x5D // Inuktitut (iu)
	SUBLANG_INUKTITUT_CANADA_LATIN = 0x02 // Canada (CA), Latin
	SUBLANG_INUKTITUT_CANADA       = 0x01 // Canada (CA), Canadian Syllabics

	LANG_IRISH            = 0x3C // Irish (ga)
	SUBLANG_IRISH_IRELAND = 0x02 // Ireland (IE)

	LANG_ITALIAN          = 0x10 // Italian (it)
	SUBLANG_ITALIAN       = 0x01 // Italy (IT)
	SUBLANG_ITALIAN_SWISS = 0x02 // Switzerland (CH)

	LANG_JAPANESE          = 0x11 // Japanese (ja)
	SUBLANG_JAPANESE_JAPAN = 0x01 // Japan (JP)

	LANG_KANNADA          = 0x4B // Kannada (kn)
	SUBLANG_KANNADA_INDIA = 0x01 // India (IN)

	LANG_KASHMIRI          = 0x60 // (reserved)
	SUBLANG_KASHMIRI_INDIA = 0x02 // (reserved)
	SUBLANG_KASHMIRI_SASIA = 0x02 // (reserved)

	LANG_KAZAK               = 0x3F // Kazakh (kk)
	SUBLANG_KAZAK_KAZAKHSTAN = 0x01 // Kazakhstan (KZ)

	LANG_KHMER             = 0x53 // Khmer (kh)
	SUBLANG_KHMER_CAMBODIA = 0x01 // Cambodia (KH)

	LANG_KICHE              = 0x86 // K'iche (qut)
	SUBLANG_KICHE_GUATEMALA = 0x01 // Guatemala (GT)

	LANG_KINYARWANDA           = 0x87 // Kinyarwanda (rw)
	SUBLANG_KINYARWANDA_RWANDA = 0x01 // Rwanda (RW)

	LANG_KONKANI          = 0x57 // Konkani (kok)
	SUBLANG_KONKANI_INDIA = 0x01 // India (IN)

	LANG_KOREAN    = 0x12 // Korean (ko)
	SUBLANG_KOREAN = 0x01 // Korea (KR)

	LANG_KYRGYZ               = 0x40 // Kyrgyz (ky)
	SUBLANG_KYRGYZ_KYRGYZSTAN = 0x01 // Kyrgyzstan (KG)

	LANG_LAO        = 0x54 // Lao (lo)
	SUBLANG_LAO_LAO = 0x01 // Lao PDR (LA)

	LANG_LATVIAN           = 0x26 // Latvian (lv)
	SUBLANG_LATVIAN_LATVIA = 0x01 // Latvia (LV)

	LANG_LITHUANIAN              = 0x27 // Lithuanian (lt)
	SUBLANG_LITHUANIAN_LITHUANIA = 0x01 // Lithuanian (LT)

	LANG_LOWER_SORBIAN            = 0x2E // Lower Sorbian (dsb)
	SUBLANG_LOWER_SORBIAN_GERMANY = 0x02 // Germany (DE)

	LANG_LUXEMBOURGISH               = 0x6E // Luxembourgish (lb)
	SUBLANG_LUXEMBOURGISH_LUXEMBOURG = 0x01 // Luxembourg (LU)

	LANG_MACEDONIAN              = 0x2F // Macedonian (mk)
	SUBLANG_MACEDONIAN_MACEDONIA = 0x01 // Macedonia (FYROM) (MK)

	LANG_MALAY                      = 0x3E // Malay (ms)
	SUBLANG_MALAY_BRUNEI_DARUSSALAM = 0x02 // Brunei Darassalam (BN)
	SUBLANG_MALAY_MALAYSIA          = 0x01 // Malaysia (MY)

	LANG_MALAYALAM          = 0x4C // Malayalam (ml)
	SUBLANG_MALAYALAM_INDIA = 0x01 // India (IN)

	LANG_MALTESE          = 0x3A // Maltese (mt)
	SUBLANG_MALTESE_MALTA = 0x01 // Malta (MT)

	LANG_MANIPURI = 0x58 // (reserved)

	LANG_MAORI                = 0x81 // Maori (mi)
	SUBLANG_MAORI_NEW_ZEALAND = 0x01 // New Zealand (NZ)

	LANG_MAPUDUNGUN          = 0x7A // Mapudungun (arn)
	SUBLANG_MAPUDUNGUN_CHILE = 0x01 // Chile (CL)

	LANG_MARATHI          = 0x4E // Marathi (mr)
	SUBLANG_MARATHI_INDIA = 0x01 // India (IN)

	LANG_MOHAWK           = 0x7C // Mohawk (moh)
	SUBLANG_MOHAWK_MOHAWK = 0x01 // Canada (CA)

	LANG_MONGOLIAN                     = 0x50 // Mongolian (mn)
	SUBLANG_MONGOLIAN_CYRILLIC_MONGOLI = 0x01 // Mongolia, Cyrillic (MN)
	SUBLANG_MONGOLIAN_PRC              = 0x02 // Mongolia, Mong (MN)

	LANG_NEPALI          = 0x61 // Nepali (ne)
	SUBLANG_NEPALI_NEPAL = 0x01 // Nepal (NP)
	SUBLANG_NEPALI_INDIA = 0x02 // India (IN)

	LANG_NORWEGIAN            = 0x14 // Norwegian (no)
	SUBLANG_NORWEGIAN_BOKMAL  = 0x01 // Bokmål, Norway (NO)
	SUBLANG_NORWEGIAN_NYNORSK = 0x02 // Nynorsk, Norway (NO)

	LANG_OCCITAN           = 0x82 // Occitan (oc)
	SUBLANG_OCCITAN_FRANCE = 0x01 // France (FR)

	LANG_ORIYA          = 0x48 // Odia (or)
	SUBLANG_ORIYA_INDIA = 0x01 // India (IN)

	LANG_PASHTO                = 0x63 // Pashto (ps)
	SUBLANG_PASHTO_AFGHANISTAN = 0x01 // Afghanistan (AF)

	LANG_PERSIAN         = 0x29 // Persian (fa)
	SUBLANG_PERSIAN_IRAN = 0x01 // Iran (IR)

	LANG_POLISH           = 0x15 // Polish (pl)
	SUBLANG_POLISH_POLAND = 0x01 // Poland (PL)

	LANG_PORTUGUESE              = 0x16 // Portuguese (pt)
	SUBLANG_PORTUGUESE_BRAZILIAN = 0x01 // Brazil (BR)
	SUBLANG_PORTUGUESE           = 0x02 // Portugal (PT)

	LANG_PULAR            = 0x67 // Pular (ff)
	SUBLANG_PULAR_SENEGAL = 0x02 // Senegal (SN)

	LANG_PUNJABI             = 0x46 // Punjabi (pa)
	SUBLANG_PUNJABI_INDIA    = 0x01 // India, Gurmukhi script (IN)
	SUBLANG_PUNJABI_PAKISTAN = 0x02 // Pakistan, Arabic script(PK)

	LANG_QUECHUA            = 0x6B // Quechua (quz)
	SUBLANG_QUECHUA_BOLIVIA = 0x01 // Bolivia (BO)
	SUBLANG_QUECHUA_ECUADOR = 0x02 // Ecuador (EC)
	SUBLANG_QUECHUA_PERU    = 0x03 // Peru (PE)

	LANG_ROMANIAN            = 0x18 // Romanian (ro)
	SUBLANG_ROMANIAN_ROMANIA = 0x01 // Romania (RO)

	LANG_ROMANSH                = 0x17 // Romansh (rm)
	SUBLANG_ROMANSH_SWITZERLAND = 0x01 // Switzerland (CH)

	LANG_RUSSIAN           = 0x19 // Russian (ru)
	SUBLANG_RUSSIAN_RUSSIA = 0x01 // Russia (RU)

	LANG_SAKHA           = 0x85 // Sakha (sah)
	SUBLANG_SAKHA_RUSSIA = 0x01 // Russia (RU)

	LANG_SAMI                     = 0x3B // Sami (smn)
	SUBLANG_SAMI_INARI_FINLAND    = 0x09 //         Inari, Finland (FI)
	SUBLANG_SAMI_LULE_NORWAY      = 0x04 // Sami (smj)  Lule, Norway (NO)
	SUBLANG_SAMI_LULE_SWEDEN      = 0x05 //         Lule, Sweden (SE)
	SUBLANG_SAMI_NORTHERN_FINLAND = 0x03 // Sami (se)   Northern, Finland (FI)
	SUBLANG_SAMI_NORTHERN_NORWAY  = 0x01 //         Northern, Norway (NO)
	SUBLANG_SAMI_NORTHERN_SWEDEN  = 0x02 //         Northern, Sweden (SE)
	SUBLANG_SAMI_SKOLT_FINLAND    = 0x08 // Sami (sms)  Skolt, Finland (FI)
	SUBLANG_SAMI_SOUTHERN_NORWA0Y = 0x06 // Sami (sma)  Southern, Norway (NO)
	SUBLANG_SAMI_SOUTHERN_SWEDEN  = 0x07 //         Southern, Sweden (SE)

	LANG_SANSKRIT          = 0x4F // Sanskrit (sa)
	SUBLANG_SANSKRIT_INDIA = 0x01 // India (IN)

	LANG_SERBIAN_NEUTRAL                        = 0x7C1A // Serbian (sr) - Neutral
	LANG_SERBIAN                                = 0x1A   // Serbian (sr)
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_CYRILLIC = 0x07   // Bosnia and Herzegovina, Cyrillic (BA)
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_LATIN    = 0x06   // Bosnia and Herzegovina, Latin (BA)
	SUBLANG_SERBIAN_CROATIA                     = 0x01   // Croatia (HR)
	SUBLANG_SERBIAN_CYRILLIC                    = 0x03   // Serbia and Montenegro (former), Cyrillic (CS)
	SUBLANG_SERBIAN_LATIN                       = 0x02   // Serbia and Montenegro (former), Latin (CS)

	LANG_SOTHO                          = 0x6C // Sesotho sa Leboa (nso)
	SUBLANG_SOTHO_NORTHERN_SOUTH_AFRICA = 0x01 // South Africa (ZA)

	LANG_TSWANA                 = 0x32 // Setswana / Tswana (tn)
	SUBLANG_TSWANA_BOTSWANA     = 0x02 // Botswana (BW)
	SUBLANG_TSWANA_SOUTH_AFRICA = 0x01 // South Africa (ZA)

	LANG_SINDHI                = 0x59 // Sindhi (sd)
	SUBLANG_SINDHI_AFGHANISTAN = 0x02 // (reserved)
	SUBLANG_SINDHI_INDIA       = 0x01 // (reserved)
	SUBLANG_SINDHI_PAKISTAN    = 0x02 // Pakistan (PK)

	LANG_SINHALESE              = 0x5B // Sinhala (si)
	SUBLANG_SINHALESE_SRI_LANKA = 0x01 // Sri Lanka (LK)

	LANG_SLOVAK             = 0x1B // Slovak (sk)
	SUBLANG_SLOVAK_SLOVAKIA = 0x01 // Slovakia (SK)

	LANG_SLOVENIAN             = 0x24 // Slovenian (sl)
	SUBLANG_SLOVENIAN_SLOVENIA = 0x01 // Slovenia (SI)

	LANG_SPANISH                       = 0x0A // Spanish (es) - 0x0A
	SUBLANG_SPANISH_ARGENTINA          = 0x0B // Argentina (AR)
	SUBLANG_SPANISH_BOLIVIA            = 0x10 // Bolivia (BO)
	SUBLANG_SPANISH_CHILE              = 0x0D // Chile (CL)
	SUBLANG_SPANISH_COLOMBIA           = 0x09 // Colombia (CO)
	SUBLANG_SPANISH_COSTA_RICA         = 0x05 // Costa Rica (CR)
	SUBLANG_SPANISH_DOMINICAN_REPUBLIC = 0x07 // Dominican Republic (DO)
	SUBLANG_SPANISH_ECUADOR            = 0x0C // Ecuador (EC)
	SUBLANG_SPANISH_EL_SALVADOR        = 0x11 // El Salvador (SV)
	SUBLANG_SPANISH_GUATEMALA          = 0x04 // Guatemala (GT)
	SUBLANG_SPANISH_HONDURAS           = 0x12 // Honduras (HN)
	SUBLANG_SPANISH_MEXICAN            = 0x02 // Mexico (MX)
	SUBLANG_SPANISH_NICARAGUA          = 0x13 // Nicaragua (NI)
	SUBLANG_SPANISH_PANAMA             = 0x06 // Panama (PA)
	SUBLANG_SPANISH_PARAGUAY           = 0x0F // Paraguay (PY)
	SUBLANG_SPANISH_PERU               = 0x0A // Peru (PE)
	SUBLANG_SPANISH_PUERTO_RICO        = 0x14 // Puerto Rico (PR)
	SUBLANG_SPANISH_MODERN             = 0x03 // Spain, Modern Sort (ES)
	SUBLANG_SPANISH                    = 0x01 // Spain, Traditional Sort (ES)
	SUBLANG_SPANISH_US                 = 0x15 // United States (US)
	SUBLANG_SPANISH_URUGUAY            = 0x0E // Uruguay (UY)
	SUBLANG_SPANISH_VENEZUELA          = 0x08 // Venezuela (VE)

	LANG_SWAHILI    = 0x41 // Swahili (sw)
	SUBLANG_SWAHILI = 0x01 // Kenya (KE)

	LANG_SWEDISH            = 0x1D // Swedish (sv)
	SUBLANG_SWEDISH_FINLAND = 0x02 // Finland (FI)
	SUBLANG_SWEDISH         = 0x01 // Sweden (SE)
	SUBLANG_SWEDISH_SWEDEN  = 0x01 // Sweden (SE)

	LANG_SYRIAC    = 0x5A // Syriac (syr)
	SUBLANG_SYRIAC = 0x01 // Syria (SY)

	LANG_TAJIK               = 0x28 // Tajik (tg)
	SUBLANG_TAJIK_TAJIKISTAN = 0x01 // Tajikistan, Cyrillic (TJ)

	LANG_TAMAZIGHT                  = 0x5F // Tamazight (tzm)
	SUBLANG_TAMAZIGHT_ALGERIA_LATIN = 0x02 // Algeria, Latin (DZ)

	LANG_TAMIL              = 0x49 // Tamil (ta)
	SUBLANG_TAMIL_INDIA     = 0x01 // India (IN)
	SUBLANG_TAMIL_SRI_LANKA = 0x02 // Sri Lanka (LK)

	LANG_TATAR           = 0x44 // Tatar (tt)
	SUBLANG_TATAR_RUSSIA = 0x01 // Russia (RU)

	LANG_TELUGU          = 0x4A // Telugu (te)
	SUBLANG_TELUGU_INDIA = 0x01 // India (IN)

	LANG_THAI             = 0x1E // Thai (th)
	SUBLANG_THAI_THAILAND = 0x01 // Thailand (TH)

	LANG_TIBETAN        = 0x51 // Tibetan (bo)
	SUBLANG_TIBETAN_PRC = 0x01 // PRC (CN)

	LANG_TIGRINYA             = 0x73 // Tigrinya (ti)
	SUBLANG_TIGRINYA_ERITREA  = 0x02 // Eritrea (ER)
	SUBLANG_TIGRINYA_ETHIOPIA = 0x01 // Ethiopia (ET)

	LANG_TIGRIGNA            = 0x73 // (reserved)
	SUBLANG_TIGRIGNA_ERITREA = 0x02 // (reserved)

	LANG_TURKISH           = 0x1F // Turkish (tr)
	SUBLANG_TURKISH_TURKEY = 0x01 // Turkey (TR)

	LANG_TURKMEN                 = 0x42 // Turkmen (tk)
	SUBLANG_TURKMEN_TURKMENISTAN = 0x01 // Turkmenistan (TM)

	LANG_UKRAINIAN            = 0x22 // Ukrainian (uk)
	SUBLANG_UKRAINIAN_UKRAINE = 0x01 // Ukraine (UA)

	LANG_UPPER_SORBIAN            = 0x2E // Upper Sorbian (hsb)
	SUBLANG_UPPER_SORBIAN_GERMANY = 0x01 // Germany (DE)

	LANG_URDU             = 0x20 // Urdu (ur)
	SUBLANG_URDU_INDIA    = 0x02 // (reserved)
	SUBLANG_URDU_PAKISTAN = 0x01 // Pakistan (PK)

	LANG_UIGHUR        = 0x80 // Uyghur (ug)
	SUBLANG_UIGHUR_PRC = 0x01 // PRC (CN)

	LANG_UZBEK             = 0x43 // Uzbek (uz)
	SUBLANG_UZBEK_CYRILLIC = 0x02 // Uzbekistan, Cyrillic (UZ)
	SUBLANG_UZBEK_LATIN    = 0x01 // Uzbekistan, Latin (UZ)

	LANG_VALENCIAN             = 0x03 // Valencian (ca)
	SUBLANG_VALENCIAN_VALENCIA = 0x02 // Valencia (ES-Valencia)

	LANG_VIETNAMESE            = 0x2A // Vietnamese (vi)
	SUBLANG_VIETNAMESE_VIETNAM = 0x01 // Vietnam (VN)

	LANG_WELSH                   = 0x52 // Welsh (cy)
	SUBLANG_WELSH_UNITED_KINGDOM = 0x01 // United Kingdom (GB)

	LANG_WOLOF            = 0x88 // Wolof (wo)
	SUBLANG_WOLOF_SENEGAL = 0x01 // Senegal (SN)

	LANG_XHOSA                 = 0x34 // isiXhosa (xh)
	SUBLANG_XHOSA_SOUTH_AFRICA = 0x01 // South Africa (ZA)

	LANG_YI        = 0x78 // Yi (ii)
	SUBLANG_YI_PRC = 0x01 // PRC (CN)

	LANG_YORUBA            = 0x6A // Yoruba (yo)
	SUBLANG_YORUBA_NIGERIA = 0x01 // Nigeria (NG)

	LANG_ZULU                 = 0x35 // isiZulu (zu)
	SUBLANG_ZULU_SOUTH_AFRICA = 0x01 // South Africa (ZA)
)

// MakeLangID https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-makelangid
// 從這個文檔，可以得知MakeLangID返回一個16個bit，前面6碼為subLang之後才是primaryLang: https://learn.microsoft.com/en-us/windows/win32/intl/language-identifiers?redirectedfrom=MSDN
// https://www.freepascal.org/docs-html/rtl/system/makelangid.html
// ★ https://learn.microsoft.com/en-us/previous-versions/windows/embedded/ms906225(v=msdn.10) 此連結底下有: #define MAKELANGID(p, s) ((((WORD) (s)) << 10) | (WORD) (p))
// AHK: https://hotkeyit.github.io/v2/docs/commands/MAKELANGID.htm
//
// 範例:
// https://renenyffenegger.ch/notes/Windows/development/Internationalization/language
// 德語: Primary: LANG_GERMAN, subLang: {DE 1, CH 2, AT 3, LU 4, LI 5}
// 所以:
// 1031 = 1*1024 + 7 for German (Germany)
// 2055 = 2*1024 + 7 for German (Switzerland)
// 3079 = 3*1024 + 7 for German (Austria)
// 4103 = 4*1024 + 7 for German (Luxembourg)
// 5127 = 5*1024 + 7 for German (Liechtenstein) = 5<<10 | 7 (用位元運算比加減快)
func MakeLangID(primaryLanguage, subLanguage uint16) uint16 {
	return subLanguage<<10 | primaryLanguage
}

// File Access Rights Constants
// https://learn.microsoft.com/en-us/windows/win32/fileio/file-access-rights-constants
const (
	FILE_ADD_FILE             = 2
	FILE_ADD_SUBDIRECTORY     = 4
	FILE_APPEND_DATA          = 4
	FILE_CREATE_PIPE_INSTANCE = 4
	FILE_DELETE_CHILD         = 64
	FILE_EXECUTE              = 32
	FILE_LIST_DIRECTORY       = 1
	FILE_READ_ATTRIBUTES      = 128
	FILE_READ_DATA            = 1
	FILE_READ_EA              = 8
	FILE_TRAVERSE             = 32
	FILE_WRITE_ATTRIBUTES     = 256
	FILE_WRITE_DATA           = 2
	FILE_WRITE_EA             = 16
)
