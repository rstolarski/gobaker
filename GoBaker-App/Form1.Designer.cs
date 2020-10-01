﻿namespace GoBaker_App
{
    partial class Form1
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.button1 = new System.Windows.Forms.Button();
            this.button2 = new System.Windows.Forms.Button();
            this.button3 = new System.Windows.Forms.Button();
            this.button4 = new System.Windows.Forms.Button();
            this.renderSizeBox = new System.Windows.Forms.TextBox();
            this.label1 = new System.Windows.Forms.Label();
            this.button5 = new System.Windows.Forms.Button();
            this.checkBox1 = new System.Windows.Forms.CheckBox();
            this.label2 = new System.Windows.Forms.Label();
            this.maxFrontTextBox = new System.Windows.Forms.TextBox();
            this.label3 = new System.Windows.Forms.Label();
            this.maxRearTextBox = new System.Windows.Forms.TextBox();
            this.SuspendLayout();
            // 
            // button1
            // 
            this.button1.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F);
            this.button1.Location = new System.Drawing.Point(58, 71);
            this.button1.Name = "button1";
            this.button1.Size = new System.Drawing.Size(178, 41);
            this.button1.TabIndex = 0;
            this.button1.Text = "LowPoly OBJ";
            this.button1.UseVisualStyleBackColor = true;
            this.button1.Click += new System.EventHandler(this.button1_Click);
            // 
            // button2
            // 
            this.button2.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F);
            this.button2.Location = new System.Drawing.Point(58, 118);
            this.button2.Name = "button2";
            this.button2.Size = new System.Drawing.Size(178, 41);
            this.button2.TabIndex = 1;
            this.button2.Text = "HighPoly OBJ";
            this.button2.UseVisualStyleBackColor = true;
            this.button2.Click += new System.EventHandler(this.button2_Click);
            // 
            // button3
            // 
            this.button3.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F);
            this.button3.Location = new System.Drawing.Point(58, 165);
            this.button3.Name = "button3";
            this.button3.Size = new System.Drawing.Size(178, 41);
            this.button3.TabIndex = 2;
            this.button3.Text = "HighPoly PLY";
            this.button3.UseVisualStyleBackColor = true;
            this.button3.Click += new System.EventHandler(this.button3_Click);
            // 
            // button4
            // 
            this.button4.Font = new System.Drawing.Font("Microsoft Sans Serif", 27.75F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(238)));
            this.button4.Location = new System.Drawing.Point(297, 165);
            this.button4.Name = "button4";
            this.button4.Size = new System.Drawing.Size(160, 88);
            this.button4.TabIndex = 3;
            this.button4.Text = "BAKE";
            this.button4.UseVisualStyleBackColor = true;
            this.button4.Click += new System.EventHandler(this.button4_Click);
            // 
            // renderSizeBox
            // 
            this.renderSizeBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F);
            this.renderSizeBox.Location = new System.Drawing.Point(144, 26);
            this.renderSizeBox.Name = "renderSizeBox";
            this.renderSizeBox.Size = new System.Drawing.Size(109, 26);
            this.renderSizeBox.TabIndex = 4;
            this.renderSizeBox.Text = "512";
            this.renderSizeBox.TextChanged += new System.EventHandler(this.renderSizeBox_TextChanged);
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F);
            this.label1.Location = new System.Drawing.Point(40, 29);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(98, 20);
            this.label1.TabIndex = 5;
            this.label1.Text = "Render size:";
            this.label1.Click += new System.EventHandler(this.label1_Click);
            // 
            // button5
            // 
            this.button5.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F);
            this.button5.Location = new System.Drawing.Point(58, 212);
            this.button5.Name = "button5";
            this.button5.Size = new System.Drawing.Size(178, 41);
            this.button5.TabIndex = 7;
            this.button5.Text = "Select Output";
            this.button5.UseVisualStyleBackColor = true;
            this.button5.Click += new System.EventHandler(this.button5_Click);
            // 
            // checkBox1
            // 
            this.checkBox1.AutoSize = true;
            this.checkBox1.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F);
            this.checkBox1.Location = new System.Drawing.Point(310, 28);
            this.checkBox1.Name = "checkBox1";
            this.checkBox1.Size = new System.Drawing.Size(132, 24);
            this.checkBox1.TabIndex = 8;
            this.checkBox1.Text = "Read ID Map?";
            this.checkBox1.UseVisualStyleBackColor = true;
            this.checkBox1.CheckedChanged += new System.EventHandler(this.checkBox1_CheckedChanged);
            // 
            // label2
            // 
            this.label2.AutoSize = true;
            this.label2.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F);
            this.label2.Location = new System.Drawing.Point(269, 74);
            this.label2.Name = "label2";
            this.label2.Size = new System.Drawing.Size(163, 20);
            this.label2.TabIndex = 10;
            this.label2.Text = "Max Frontal Distance:";
            this.label2.Click += new System.EventHandler(this.label2_Click_1);
            // 
            // maxFrontTextBox
            // 
            this.maxFrontTextBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F);
            this.maxFrontTextBox.Location = new System.Drawing.Point(447, 71);
            this.maxFrontTextBox.Name = "maxFrontTextBox";
            this.maxFrontTextBox.Size = new System.Drawing.Size(109, 26);
            this.maxFrontTextBox.TabIndex = 9;
            this.maxFrontTextBox.Text = "3.0";
            this.maxFrontTextBox.TextChanged += new System.EventHandler(this.maxFrontTextBox_TextChanged);
            // 
            // label3
            // 
            this.label3.AutoSize = true;
            this.label3.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F);
            this.label3.Location = new System.Drawing.Point(269, 106);
            this.label3.Name = "label3";
            this.label3.Size = new System.Drawing.Size(148, 20);
            this.label3.TabIndex = 12;
            this.label3.Text = "Max Rear Distance:";
            // 
            // maxRearTextBox
            // 
            this.maxRearTextBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F);
            this.maxRearTextBox.Location = new System.Drawing.Point(447, 103);
            this.maxRearTextBox.Name = "maxRearTextBox";
            this.maxRearTextBox.Size = new System.Drawing.Size(109, 26);
            this.maxRearTextBox.TabIndex = 11;
            this.maxRearTextBox.Text = "3.0";
            this.maxRearTextBox.TextChanged += new System.EventHandler(this.maxRearTextBox_TextChanged);
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(594, 287);
            this.Controls.Add(this.label3);
            this.Controls.Add(this.maxRearTextBox);
            this.Controls.Add(this.label2);
            this.Controls.Add(this.maxFrontTextBox);
            this.Controls.Add(this.checkBox1);
            this.Controls.Add(this.button5);
            this.Controls.Add(this.label1);
            this.Controls.Add(this.renderSizeBox);
            this.Controls.Add(this.button4);
            this.Controls.Add(this.button3);
            this.Controls.Add(this.button2);
            this.Controls.Add(this.button1);
            this.Name = "Form1";
            this.Text = "Form1";
            this.Load += new System.EventHandler(this.Form1_Load);
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.Button button1;
        private System.Windows.Forms.Button button2;
        private System.Windows.Forms.Button button3;
        private System.Windows.Forms.Button button4;
        private System.Windows.Forms.TextBox renderSizeBox;
        private System.Windows.Forms.Label label1;
        private System.Windows.Forms.Button button5;
        private System.Windows.Forms.CheckBox checkBox1;
        private System.Windows.Forms.Label label2;
        private System.Windows.Forms.TextBox maxFrontTextBox;
        private System.Windows.Forms.Label label3;
        private System.Windows.Forms.TextBox maxRearTextBox;
    }
}

