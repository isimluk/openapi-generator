/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Org.OpenAPITools.Client.OpenAPIDateConverter;
using OpenAPIClientUtils = Org.OpenAPITools.Client.ClientUtils;

namespace Org.OpenAPITools.Model
{
    /// <summary>
    /// Drawing
    /// </summary>
    [DataContract(Name = "Drawing")]
    public partial class Drawing : Dictionary<String, Fruit>, IEquatable<Drawing>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="Drawing" /> class.
        /// </summary>
        /// <param name="mainShape">mainShape.</param>
        /// <param name="shapeOrNull">shapeOrNull.</param>
        /// <param name="nullableShape">nullableShape.</param>
        /// <param name="shapes">shapes.</param>
        public Drawing(Shape mainShape = default(Shape), ShapeOrNull shapeOrNull = default(ShapeOrNull), NullableShape nullableShape = default(NullableShape), List<Shape> shapes = default(List<Shape>)) : base()
        {
            this._MainShape = mainShape;
            this._ShapeOrNull = shapeOrNull;
            this._NullableShape = nullableShape;
            this._Shapes = shapes;
        }

        /// <summary>
        /// Gets or Sets MainShape
        /// </summary>
        [DataMember(Name = "mainShape", EmitDefaultValue = false)]
        public Shape MainShape
        {
            get{ return _MainShape;}
            set
            {
                _MainShape = value;
                _flagMainShape = true;
            }
        }
        private Shape _MainShape;
        private bool _flagMainShape;

        /// <summary>
        /// Returns false as MainShape should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeMainShape()
        {
            return _flagMainShape;
        }
        /// <summary>
        /// Gets or Sets ShapeOrNull
        /// </summary>
        [DataMember(Name = "shapeOrNull", EmitDefaultValue = false)]
        public ShapeOrNull ShapeOrNull
        {
            get{ return _ShapeOrNull;}
            set
            {
                _ShapeOrNull = value;
                _flagShapeOrNull = true;
            }
        }
        private ShapeOrNull _ShapeOrNull;
        private bool _flagShapeOrNull;

        /// <summary>
        /// Returns false as ShapeOrNull should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeShapeOrNull()
        {
            return _flagShapeOrNull;
        }
        /// <summary>
        /// Gets or Sets NullableShape
        /// </summary>
        [DataMember(Name = "nullableShape", EmitDefaultValue = true)]
        public NullableShape NullableShape
        {
            get{ return _NullableShape;}
            set
            {
                _NullableShape = value;
                _flagNullableShape = true;
            }
        }
        private NullableShape _NullableShape;
        private bool _flagNullableShape;

        /// <summary>
        /// Returns false as NullableShape should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeNullableShape()
        {
            return _flagNullableShape;
        }
        /// <summary>
        /// Gets or Sets Shapes
        /// </summary>
        [DataMember(Name = "shapes", EmitDefaultValue = false)]
        public List<Shape> Shapes
        {
            get{ return _Shapes;}
            set
            {
                _Shapes = value;
                _flagShapes = true;
            }
        }
        private List<Shape> _Shapes;
        private bool _flagShapes;

        /// <summary>
        /// Returns false as Shapes should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeShapes()
        {
            return _flagShapes;
        }
        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class Drawing {\n");
            sb.Append("  ").Append(base.ToString().Replace("\n", "\n  ")).Append("\n");
            sb.Append("  MainShape: ").Append(MainShape).Append("\n");
            sb.Append("  ShapeOrNull: ").Append(ShapeOrNull).Append("\n");
            sb.Append("  NullableShape: ").Append(NullableShape).Append("\n");
            sb.Append("  Shapes: ").Append(Shapes).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input as Drawing).AreEqual;
        }

        /// <summary>
        /// Returns true if Drawing instances are equal
        /// </summary>
        /// <param name="input">Instance of Drawing to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(Drawing input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input).AreEqual;
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = base.GetHashCode();
                if (this.MainShape != null)
                    hashCode = hashCode * 59 + this.MainShape.GetHashCode();
                if (this.ShapeOrNull != null)
                    hashCode = hashCode * 59 + this.ShapeOrNull.GetHashCode();
                if (this.NullableShape != null)
                    hashCode = hashCode * 59 + this.NullableShape.GetHashCode();
                if (this.Shapes != null)
                    hashCode = hashCode * 59 + this.Shapes.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
